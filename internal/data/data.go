package data

import (
	"database/sql"
	"fmt"
	"html"
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

func getChannelFromSource(c *pb.Channel) error {
	if c.Id == "" || c.Id == "cid" {
		return errors.New("nil or wrong channel id: " + c.Id)
	}
	// https://www.youtube.com/channel/UCMUnInmOkrWN4gof9KlhNmQ/videos
	u, err := url.Parse("https://www.youtube.com/channel/" + c.Id + "/videos")
	raw, _, err := exhtml.GetRawAndDoc(u, 1*time.Minute)
	if err != nil {
		return err
	}
	// get channel name
	// c.Name = exhtml.ElementsByTag(doc, "title")[0].FirstChild.Data
	// c.Name = strings.Replace(c.Name, " - YouTube", "", -1)
	re := regexp.MustCompile(`<title>(.*) - YouTube</title>`)
	rs := re.FindAllSubmatch(raw, -1)
	if len(rs) == 0 {
		return errors.New("getChannelFromSource get channel name match nothing:" + c.Id)
	}
	c.Name = html.UnescapeString(string(rs[0][1]))
	// get vids
	re = regexp.MustCompile(`"gridVideoRenderer":{"videoId":"(.*?)","thumbnail":{"thumbnails"`)
	rs = re.FindAllSubmatch(raw, -1)
	if len(rs) == 0 {
		return errors.New("getChannelFromSource get vids match nothing")
	}
	for _, r := range rs {
		c.Vids = append(c.Vids, string(r[1]))
	}
	return nil
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
	return getVideo(dc, v)
}

// getVideo get video info if it's Id is currect
// if video info not in db, it will obtain cid by api source and others by youtube pkg
// then INSERT OR UPDATE TABLES: videos and channels.
func getVideo(dc *sql.DB, v *pb.Video) (*pb.Video, error) {
	err := db.SelectVideoByVid(dc, v)
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

// getCid will get cid by vid from db first, then from source of api
func getCid(dc *sql.DB, vid string) (string, error) {
	cid, err := db.SelectCidByVid(dc, vid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return getCidFromSource(vid)
		}
		return "", err
	}
	return cid, nil
}

// TODO: get channel id and name then insert
// TODO: rename the function
func getVideoFromApi(dc *sql.DB, vid string) (*pb.Video, error) {
	v := &pb.Video{Id: vid}
	client := youtube.Client{}
	video, err := client.GetVideo("https://www.youtube.com/watch?v=" + v.Id)
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
	cid := c.Id
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
// 3. insert or update tables: videos and channels
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

// GetChannel query channel info
// If nothing got from database, get video ids and channel info from source web page.
func (fr *fetcherRepo) GetChannel(c *pb.Channel) error {
	// Select name from channels
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	if err = db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Get video ids and channel info from source
			if err := getChannelFromSource(c); err != nil {
				return err
			}
			return db.InsertChannel(dc, c) // storage channel info just got
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

	if err = db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err = getChannelFromSource(c); err != nil {
				return err
			}
			return db.InsertChannel(dc, c) // storage channel info just got
		}
		return err
	}
	return nil
}
