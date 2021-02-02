package biz

import (
	"fmt"
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"log"
	"net/url"
)

type Fetcher struct {
	Entrance *url.URL
	Links    []string
}

type FetcherCase struct {
	repo FetcherRepo
}

type FetcherRepo interface {
	GetLinks(*Fetcher) ([]string, error)
	GetVideos(*Fetcher) ([]*pb.Video, error)
}

func NewFetcherCase(repo FetcherRepo) *FetcherCase {
	return &FetcherCase{repo: repo}
}

func (fc *FetcherCase) Crawl(f *Fetcher) error {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Crawl err: %v", err)
		}
	}()
	fmt.Println("fetch links from ", f.Entrance.String())
	var err error
	if f.Links, err = fc.repo.GetLinks(f); err != nil {
		return err
	}
	return nil
}
