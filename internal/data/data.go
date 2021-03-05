package data

import (
	"database/sql"
	"fmt"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
	"github.com/pkg/errors"
)

var _ biz.FetcherRepo = new(fetcherRepo)

type fetcherRepo struct{}

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
// Notice: If greedy, it will storage vids and cid to db
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
	if err = db.InsertVids(dc, c.Vids, c.Id); err != nil {
		return nil, err
	}
	return c, nil
}

// GetVids obtain videoIds by c.Id, default is ungreedy
func (fr *fetcherRepo) GetVids(c *pb.Channel) (*pb.Channel, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return getVids(dc, c, false)
}

// GetVideos get or (if greedy) storage videos info to db by videos page of the channel
func (fr *fetcherRepo) GetVideos(c *pb.Channel) ([]*pb.Video, error) {
	greedy := false // so, it will get videos by db search only
	c, err := fr.GetVids(c)
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
func getVideos(dc *sql.DB, c *pb.Channel, greedy bool) ([]*pb.Video, error) {
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
	return vs, nil
}

// GetChannel query channel info
// If nothing got from database, get video ids and channel info from source web page.
func (fr *fetcherRepo) GetChannel(c *pb.Channel) error {
	// Select name from channels
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	return getChannel(dc, c, false)
}

func getChannel(dc *sql.DB, c *pb.Channel, greedy bool) error {
	if err := db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) && greedy { // if greedy and no rows found out
			// Get video ids and channel info from source
			if err := getChannelFromSource(c); err != nil {
				return err
			}
			if err = db.InsertChannel(dc, c); err != nil { // storage channel info just got
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

func (fr *fetcherRepo) GetChannelName(c *pb.Channel) error {
	// Select name from channels
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()
	return getChannel(dc, c, false)
}

func getChannelName(dc *sql.DB, c *pb.Channel, greedy bool) error {
	if err := db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) && greedy { // if greedy and no rows found out
			if err = getChannelFromSource(c); err != nil {
				return err
			}
			return db.InsertChannel(dc, c) // storage channel info just got
		}
		return err
	}
	return nil
}

func (fr *fetcherRepo) GetChannels(cs *pb.Channels) (*pb.Channels, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	return db.SelectChannels(dc, cs)
}
