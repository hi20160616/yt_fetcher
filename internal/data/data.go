package data

import (
	"fmt"
	"net/url"
	"time"

	"github.com/hi20160616/exhtml"
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	youtube "github.com/kkdai/youtube/v2"
)

var _ biz.FetcherRepo = new(fetcherRepo)

type fetcherRepo struct{}

func (fr *fetcherRepo) GetLinks(f *biz.Fetcher) ([]string, error) {
	links, err := exhtml.ExtractLinks(f.Entrance.String())
	if err != nil {
		return nil, err
	}
	f.Links = links
	return links, nil
}

func (fr *fetcherRepo) GetVideos(f *biz.Fetcher) ([]*pb.Video, error) {
	for _, link := range f.Links {
		// fr.GetVideo()
		fmt.Println("video links: " + link)
	}
	return nil, nil
}

func (fr *fetcherRepo) GetVideo(_url string) (*pb.Video, error) {
	u, err := url.Parse(_url)
	if err != nil {
		return nil, err
	}
	videoID := u.Query().Get("v")
	client := youtube.Client{}
	video, err := client.GetVideo(videoID)
	if err != nil {
		return nil, err
	}
	v := &pb.Video{
		Id:          video.ID,
		Title:       video.Title,
		Description: video.Description,
		Author:      video.Author,
		// LastUpdated: time.Now(),
	}

	return v, nil
}
