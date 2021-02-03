package biz

import (
	"errors"
	"net/url"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

type Fetcher struct {
	Entrance *url.URL
	Links    []string
}

type FetcherCase struct {
	repo FetcherRepo
}

type FetcherRepo interface {
	GetAndSetLinks(*Fetcher) ([]string, error)
	NewVideo(string) (*pb.Video, error)
	GetVideo(*pb.Video) (*pb.Video, error)
	GetVideos(*Fetcher) ([]*pb.Video, error)
}

func NewFetcherCase(repo FetcherRepo) *FetcherCase {
	return &FetcherCase{repo: repo}
}

func (fc *FetcherCase) GetVideos(f *Fetcher) ([]*pb.Video, error) {
	videos, err := fc.repo.GetVideos(f)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (fc *FetcherCase) GetVideo(v *pb.Video) (*pb.Video, error) {
	if v.Id == "" {
		return nil, errors.New("fc.GetVideo err: video id is nil")
	}
	video, err := fc.repo.GetVideo(v)
	if err != nil {
		return nil, err
	}
	return video, nil
}

// func (fc *FetcherCase) Crawl(f *Fetcher) error {
//         defer func() {
//                 if err := recover(); err != nil {
//                         log.Printf("Crawl err: %v", err)
//                 }
//         }()
//         fmt.Println("fetch links from ", f.Entrance.String())
//         var err error
//         if f.Links, err = fc.repo.GetAndSetLinks(f); err != nil {
//                 return err
//         }
//         return nil
// }
