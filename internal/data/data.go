package data

import (
	"database/sql"
	"fmt"
	"net/url"
	"regexp"
	"time"

	"github.com/hi20160616/exhtml"
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
	youtube "github.com/kkdai/youtube/v2"
	"github.com/pkg/errors"
)

var _ biz.FetcherRepo = new(fetcherRepo)

type fetcherRepo struct{}

func NewFetcherRepo() biz.FetcherRepo {
	return &fetcherRepo{}
}

// getVidsFromSource get vids from raw html page.
func getVidsFromSource(cid string) ([]string, error) {
	vids := []string{}
	u, err := url.Parse("https://www.youtube.com/channel/" + cid + "/videos")
	raw, _, err := exhtml.GetRawAndDoc(u, 1*time.Minute)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`"gridVideoRenderer":{"videoId":"(.*?)","thumbnail":{"thumbnails"`)
	rs := re.FindAllSubmatch(raw, -1)
	for _, r := range rs {
		vids = append(vids, string(r[1]))
	}
	return vids, nil
}

// getCidFromSource get cid from youtube video api source
func getCidFromSource(vid string) (string, error) {
	video := "http://youtube.com/get_video_info?video_id=" + vid
	u, err := url.Parse(video)
	if err != nil {
		return "", err
	}
	raw, _, err := exhtml.GetRawAndDoc(u, 1*time.Minute)
	if err != nil {
		return "", err
	}
	r, err := url.QueryUnescape(string(raw))
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`","externalChannelId":"(.*?)","availableCountries":`)
	rs := re.FindAllStringSubmatch(r, -1)
	return rs[0][1], nil

}

// NewVideo make and return Video object with this id.
func (fr *fetcherRepo) NewVideo(id string) (*pb.Video, error) {
	v := &pb.Video{Vid: id}
	if v.Vid == "" {
		return nil, fmt.Errorf("NewVideo err: id is nil")
	}
	return v, nil
}

// GetVideo get video info if it's Id is currect
// if video info not in db, it will obtain cid by api source and others by youtube pkg
func (fr *fetcherRepo) GetVideo(v *pb.Video) (*pb.Video, error) {
	if v.Vid == "" {
		return nil, fmt.Errorf("GetVideo err: video id is nil, you need fr.NewVideo(id) first.")
	}

	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return getVideo(dc, v)
}

// getVideo get video info if it's Id is currect
// if video info not in db, it will obtain cid by api source and others by youtube pkg
func getVideo(dc *sql.DB, v *pb.Video) (*pb.Video, error) {
	_v, err := selectVideoFromDb(dc, v.Vid)
	if err != nil {
		// No video in db, get from api and insert to db
		if errors.Is(err, sql.ErrNoRows) {
			_v, err = getVideoFromApi(dc, v.Vid)
			if err != nil {
				return nil, err
			}
			return _v, db.InsertOrUpdate(dc, _v)
		}
		return nil, err
	}

	if _v.Title == "" { // maybe, this is a video only have vid and cid
		_v, err = getVideoFromApi(dc, v.Vid)
		if err != nil {
			return nil, err
		}
		return _v, db.InsertOrUpdate(dc, _v)
	}
	return _v, nil
}

// selectVideoFromDb select * from db.yt_fetcher.videos where id = 'vid'
func selectVideoFromDb(dc *sql.DB, vid string) (*pb.Video, error) {
	return db.SelectVideo(dc, vid)
}

// getCid will get cid by vid from db first, then from source of api
func getCid(dc *sql.DB, vid string) (string, error) {
	cid, err := db.SelectCid(dc, vid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return getCidFromSource(vid)
		}
		return "", err
	}
	return cid, nil
}

func getVideoFromApi(dc *sql.DB, vid string) (*pb.Video, error) {
	v := &pb.Video{Vid: vid}
	client := youtube.Client{}
	video, err := client.GetVideo("https://www.youtube.com/watch?v=" + v.Vid)
	if err != nil {
		return nil, err
	}
	cid, err := getCid(dc, vid)
	if err != nil {
		return nil, err
	}
	t := video.Formats.FindByQuality("medium").LastModified
	v.Title = video.Title
	v.Description = video.Description
	v.Cid = cid
	v.Cname = video.Author
	v.LastUpdated = t
	return v, nil
}

// GetVids obtain videoIds from video page of the channel just loaded.
// it will not query from db, because video page always update frequently.
// Notice: it will storage vids and cid to db
func (fr *fetcherRepo) GetVids(c *pb.Channel) (*pb.Channel, error) {
	cid := c.Cid
	vids, err := getVidsFromSource(cid)
	if err != nil {
		return nil, err
	}
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	if err = db.InsertVids(dc, vids, cid); err != nil {
		return nil, err
	}
	c.Vids = vids
	return c, nil
}

// GetVideos get and storage videos info to db by videos page of the channel
// 1. get vids from videos page every request.
// 2. for range vids and get video info by vid
// return video slice.
func (fr *fetcherRepo) GetVideos(c *pb.Channel) ([]*pb.Video, error) {
	c, err := fr.GetVids(c)
	if err != nil {
		return nil, err
	}
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	videos := []*pb.Video{}
	for _, id := range c.Vids {
		v, err := fr.NewVideo(id)
		if err != nil {
			return nil, err
		}
		v, err = getVideo(dc, v)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	return videos, nil
}

func (fr *fetcherRepo) GetChannel(c *pb.Channel) (*pb.Channel, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	// Get and set vids to channel
	vids, err := fr.GetVids(c)
	if err != nil {
		return nil, err
	}
	c.Vids = vids.Vids

	c, err = fr.GetSetCname(c)

	return db.GetChannel(dc, c)
}

func (fr *fetcherRepo) GetSetCname(c *pb.Channel) (*pb.Channel, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return db.GetSetCname(dc, c)
}
