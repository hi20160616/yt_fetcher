package data

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/hi20160616/exhtml"
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	youtube "github.com/kkdai/youtube/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.FetcherRepo = new(fetcherRepo)

type fetcherRepo struct{}

func NewFetcherRepo() biz.FetcherRepo {
	return &fetcherRepo{}
}

// ExtractVideosIds only extract videoIds at that page loaded first.
// search database first and working the channelId if not met.
func (fr *fetcherRepo) GetVideoIds(c *pb.Channel) (*pb.Channel, error) {
	u, err := url.Parse(c.Url)
	if err != nil {
		return nil, err
	}

	vids, err := selectVidsFromDb(u.Query().Get("v"))
	if err != nil {
		return c, nil
	}

	if vids == nil {
		vids, err = getVidsFromSource(u)
		if err != nil {
			return nil, err
		}
		// TODO: DB insert vids
	}
	c.VideoIds = vids
	return c, nil
}

// selectVidsFromDb select vid from db.yt_fetcher.videos where cid = 'channelId'
func selectVidsFromDb(channelId string) ([]string, error) {
	return nil, nil
}

// getVidsFromSource get vids from raw html page.
func getVidsFromSource(u *url.URL) ([]string, error) {
	vids := []string{}
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

// NewVideo make and return Video object with this id.
func (fr *fetcherRepo) NewVideo(id string) (*pb.Video, error) {
	v := &pb.Video{Id: id}
	if v.Id == "" {
		return nil, fmt.Errorf("NewVideo err: id is nil")
	}
	return v, nil
}

// GetVideo get video info if it's Id is currect
func (fr *fetcherRepo) GetVideo(v *pb.Video) (*pb.Video, error) {
	if v.Id == "" {
		return nil, fmt.Errorf("GetVideo err: video id is nil, you need fr.NewVideo(id) first.")
	}

	_v, err := selectVideoFromDb(v.Id)
	if err != nil {
		return nil, err
	}

	if _v == nil {
		_v, err = getVideoFromApi(v.Id)
		if err != nil {
			return nil, err
		}
	}
	v = _v
	return v, nil
}

// selectVideoFromDb select * from db.yt_fetcher.videos where id = 'vid'
func selectVideoFromDb(vid string) (*pb.Video, error) {
	return nil, nil
}

func getVideoFromApi(vid string) (*pb.Video, error) {
	v := &pb.Video{Id: vid}
	client := youtube.Client{}
	video, err := client.GetVideo(v.Id)
	if err != nil {
		return nil, err
	}
	t := video.Formats.FindByQuality("medium").LastModified
	tt, err := strconv.ParseInt(t[:10], 10, 64)
	if err != nil {
		return nil, err
	}
	ttt := time.Unix(tt, 0)
	v.Title = video.Title
	v.Description = video.Description
	v.Author = video.Author
	v.LastUpdated = timestamppb.New(ttt)
	return v, nil
}

func (fr *fetcherRepo) videoIdsInit(c *pb.Channel) error {
	var err error
	if len(c.VideoIds) == 0 {
		c, err = fr.GetVideoIds(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fr *fetcherRepo) GetVideos(c *pb.Channel) ([]*pb.Video, error) {
	if err := fr.videoIdsInit(c); err != nil {
		return nil, err
	}
	videos := []*pb.Video{}
	for _, id := range c.VideoIds {
		v, err := fr.NewVideo(id)
		if err != nil {
			return nil, err
		}
		v, err = fr.GetVideo(v)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	return videos, nil
}
