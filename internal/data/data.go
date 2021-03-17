package data

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
	"github.com/pkg/errors"
)

var _ biz.FetcherRepo = new(fetcherRepo)

type fetcherRepo struct {
}

func NewFetcherRepo() biz.FetcherRepo {
	return &fetcherRepo{}
}

// NewVideo make and return Video object with this id.
func (fr *fetcherRepo) NewVideo(id string) (*pb.Video, error) {
	v := &pb.Video{Id: id}
	if v.Id == "" {
		return nil, fmt.Errorf("NewVideo err: id is nil")
	}
	return v, nil
}

// SearchVideos select all videos match the keywords,
// If no vs.Keywords provided, it'll return all the videos in table of database.
func (fr *fetcherRepo) SearchVideos(vs *pb.Videos) (*pb.Videos, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	return db.SearchVideos(dc, vs, vs.Keywords...)
}

// GetVideo get video info if it's Id is currect
// if video info not in db, it will obtain cid by api source and others by youtube pkg
func (fr *fetcherRepo) GetVideo(v *pb.Video) (*pb.Video, error) {
	if v.Id == "" {
		return nil, fmt.Errorf("GetVideo err: video id is nil, you need fr.NewVideo(id) first.")
	}

	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return getVideo(dc, v, false)
}

// getVideo get video info and set v by v.Id
// greedy: If greedy and video info not in db, it will obtain cid by api source and others by youtube pkg
// then INSERT OR UPDATE TABLES: videos and channels.
func getVideo(dc *sql.DB, v *pb.Video, greedy bool) (*pb.Video, error) {
	err := db.SelectVideoByVid(dc, v)
	if !greedy {
		return v, err
	}
	if err != nil {
		// No video in db, get from api and insert to db
		if errors.Is(err, sql.ErrNoRows) {
			v, err = getVideoFromApi(dc, v.Id)
			if err != nil {
				return nil, err
			}
			return v, db.InsertOrUpdateVC(dc, v)
		}
		return nil, err
	}

	if v.Title == "" { // maybe, this is a video only have vid and cid
		v, err = getVideoFromApi(dc, v.Id)
		if err != nil {
			return nil, err
		}
		return v, db.InsertOrUpdateVC(dc, v)
	}
	return v, nil
}

// getCid will get cid by vid from db,
// greedy: If greedy and cid not in db, obtain it from source of api
func getCid(dc *sql.DB, vid string, greedy bool) (string, error) {
	cid, err := db.SelectCidByVid(dc, vid)
	if !greedy {
		return cid, err
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return getCidFromSource(vid)
		}
		return "", err
	}
	return cid, nil
}

// getVids obtain videoIds by c.Id
// If greedy and no ids in db, fetch videos' info from channel page
// -- Func Commented! Notice: If greedy, it will storage vids and cid to db --
func getVids(dc *sql.DB, c *pb.Channel, greedy bool) (*pb.Channel, error) {
	var err error
	if !greedy {
		c.Vids, err = db.SelectVidsByCid(dc, c.Id)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	c.Vids, err = getVidsFromSource(c.Id)
	if err != nil {
		return nil, err
	}
	// if err = db.InsertVids(dc, c.Vids, c.Id); err != nil {
	//         return nil, err
	// }
	return c, nil
}

// GetVids obtain videoIds by c.Id, default is ungreedy
func (fr *fetcherRepo) GetVids(c *pb.Channel, greedy bool) (*pb.Channel, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return getVids(dc, c, greedy)
}

func (fr *fetcherRepo) GetVideosFromTo(vs *pb.Videos) (*pb.Videos, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	timeStamp := func(days int) string {
		t := time.Now().AddDate(0, 0, days).UnixNano()
		return strconv.FormatInt(t, 10)[:16]
	}

	vs.After = timeStamp(-3) // 3 days ago
	vs.Before = timeStamp(0)
	return db.SelectVideosFromTo(dc, vs)
}

// GetVideos get or (if greedy) storage videos info to db by videos page of the channel
func (fr *fetcherRepo) GetVideos(c *pb.Channel, greedy bool) (*pb.Videos, error) {
	// greedy := false // so, it will get videos by db search only
	c, err := fr.GetVids(c, greedy)
	if err != nil {
		return nil, err
	}
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return getVideos(dc, c, greedy)
}

// getVideos get videos from db by c.Id
// If greedy is true, notice:
// 1. It will get vids from videos page every request.
// 2. It will insert or update tables: videos and channels
func getVideos(dc *sql.DB, c *pb.Channel, greedy bool) (*pb.Videos, error) {
	c, err := getVids(dc, c, greedy)
	if err != nil {
		return nil, err
	}

	vs := []*pb.Video{}
	for _, id := range c.Vids {
		v := &pb.Video{Id: id}
		v, err = getVideo(dc, v, greedy)
		if err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}
	return &pb.Videos{Videos: vs}, nil
}

// GetChannel query channel info
// If nothing got from database, get video ids and channel info from source web page.
func (fr *fetcherRepo) GetChannel(c *pb.Channel) (*pb.Channel, error) {
	// Select name from channels
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	return getChannel(dc, c, false)
}

// getChannel get channel from database by id,
// it will fetch from source page if only greedy is true
func getChannel(dc *sql.DB, c *pb.Channel, greedy bool) (*pb.Channel, error) {
	if err := db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) && greedy { // if greedy and no rows found out
			// Get video ids and channel info from source
			c, err := getChannelFromSource(c)
			if err != nil {
				return nil, err
			}
			if err = db.InsertChannel(dc, c); err != nil { // storage channel info just got
				return nil, errors.WithMessage(err, "getChannel error")
			}
			return c, nil
		}
		return nil, err
	}
	return c, nil
}

func (fr *fetcherRepo) GetChannelName(c *pb.Channel) (*pb.Channel, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return getChannelName(dc, c, false)
}

// getChannelName get channel name by id from database,
// it will fetch from source page if only greedy if true
func getChannelName(dc *sql.DB, c *pb.Channel, greedy bool) (*pb.Channel, error) {
	if err := db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) && greedy { // if greedy and no rows found out
			if c, err = getChannelFromSource(c); err != nil {
				return nil, err
			}
			return c, db.InsertChannel(dc, c) // storage channel info just got
		}
		return nil, err
	}
	return c, nil
}

// GetChannels get all channels from database
func (fr *fetcherRepo) GetChannels(cs *pb.Channels) (*pb.Channels, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	return db.SelectChannels(dc, cs)
}

// DelChannel delete channel by id in c
func (fr *fetcherRepo) DelChannel(c *pb.Channel) error {
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()
	return db.DelChannel(dc, c)
}

// UpdateChannels default greedy false
func (fr *fetcherRepo) UpdateChannels(cs *pb.Channels, greedy bool) error {
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	cs, err = fr.GetChannels(cs)
	if err != nil {
		return err
	}
	for _, c := range cs.Channels {
		if err = updateChannelFromSource(dc, c, greedy); err != nil {
			return err
		}
	}
	return nil
}

// updateChannelFromSource update channel by source
// greedy true: get videos from api directly
// greedy false: get videos from api if only it is not exist in table videos
func updateChannelFromSource(dc *sql.DB, c *pb.Channel, greedy bool) error {
	// Code logic
	// 1. get vids from the channel source every request
	// 2. for range vids to get and set videos
	//    2.1. if greedy, get video directly from api and InsertOrUpdateVideo
	//    2.2. if not greedy
	//         2.2.1. if vid not exist in videos, getVideoFromApi and InsertOrUpdateVideo
	//         2.2.2. if vid exist in videos, pass the loop this time.

	// must greedy here, so get Vids from source every request
	c, err := getVids(dc, c, true)
	if err != nil {
		return err
	}
	do := func(vid string) error {
		// no matter video is updated in youtube, just get info from source
		v, err := getVideoFromApi(dc, vid)
		if err != nil {
			return err
		}
		return db.InsertOrUpdateVideo(dc, v)
	}
	for _, vid := range c.Vids {
		if greedy {
			err = do(vid)
		} else {
			exist, err := db.VidExist(dc, vid)
			if err != nil {
				return err
			}
			if !exist {
				err = do(vid)
			}
		}
	}
	if err != nil {
		return err
	}
	return db.UpdateChannel(dc, c) // actualy update chananel last_updated field.
}
