package data

import (
	"errors"
	"fmt"
	"net/url"
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

func (fr *fetcherRepo) GetAndSetLinks(f *biz.Fetcher) ([]string, error) {
	// is it can fetch by youtube api or youtube v2?
	links, err := exhtml.ExtractLinks(f.Entrance.String())
	if err != nil {
		return nil, err
	}
	f.Links = links
	return links, nil
}

// NewVideo parse _url, get video id from arg `v`, make and return Video object with this id.
func (fr *fetcherRepo) NewVideo(_url string) (*pb.Video, error) {
	u, err := url.Parse(_url)
	if err != nil {
		return nil, err
	}

	v := &pb.Video{Id: u.Query().Get("v")}
	if v.Id == "" {
		return nil, errors.New("NewVideo err: no id got from url: " + _url)
	}

	return v, nil
}

// GetVideo get video info if it's Id is currect
func (fr *fetcherRepo) GetVideo(v *pb.Video) (*pb.Video, error) {
	if v.Id == "" {
		return nil, errors.New("GetVideo err: video id is nil.")
	}
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

func (fr *fetcherRepo) GetVideos(f *biz.Fetcher) ([]*pb.Video, error) {
	videos := []*pb.Video{}
	for _, link := range f.Links {
		v, err := fr.NewVideo(link)
		if err != nil {
			return nil, err
		}
		v, err = fr.GetVideo(v)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
		fmt.Println("video links: " + link)
	}
	return videos, nil
}
