package data

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
	"github.com/pkg/errors"
)

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
// then INSERT OR UPDATE TABLES: videos, thumbnails and channels.
// TODO: need pass test for thumbnails be greedy
func getVideo(dc *sql.DB, v *pb.Video, greedy bool) (*pb.Video, error) {
	v, err := db.SelectVideoByVid(dc, v)
	if err != nil {
		return nil, err
	}
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
			return v, db.Insert2(dc, v)
		}
		return nil, err
	}

	if v.Title == "" { // maybe, this is a video only have vid and cid
		v, err = getVideoFromApi(dc, v.Id)
		if err != nil {
			return nil, err
		}
		return v, db.Insert2(dc, v)
	}
	return v, nil
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
			if errors.Is(err, ErrIgnoreVideoOnPurpose) {
				log.Println(err)
				continue
			}
			return nil, err
		}
		vs = append(vs, v)
	}
	return &pb.Videos{Videos: vs}, nil
}

// GetVideosIn24H get recent 24 hours videos in tables
func (fr *fetcherRepo) GetVideosIn24H(vs *pb.Videos) (*pb.Videos, error) {
	return fr.GetVideosFromTo(vs)
}

// GetVideosFromTo get videos from vs.After to vs.Before
// if vs.After and vs.Before are all nil, get videos 24 hours age.
func (fr *fetcherRepo) GetVideosFromTo(vs *pb.Videos) (*pb.Videos, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	timeStamp := func(minutes int) string {
		t := time.Now().Add(time.Duration(minutes) * time.Minute).UnixNano()
		return strconv.FormatInt(t, 10)[:16]
	}

	if vs.After == "" {
		vs.After = timeStamp(-1 * 24 * 60) // 1 days ago
	}
	if vs.Before == "" {
		vs.Before = timeStamp(0)
	}

	return db.SelectVideosFromTo(dc, vs)
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
