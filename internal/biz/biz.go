package biz

import (
	"errors"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

type FetcherCase struct {
	repo FetcherRepo
}

type FetcherRepo interface {
	NewVideo(string) (*pb.Video, error)
	GetVideo(*pb.Video) (*pb.Video, error)
	GetVids(*pb.Channel) (*pb.Channel, error)
	GetVideos(*pb.Channel) ([]*pb.Video, error)
}

func NewFetcherCase(repo FetcherRepo) *FetcherCase {
	return &FetcherCase{repo: repo}
}

func (fc *FetcherCase) GetVideoIds(c *pb.Channel) (*pb.Channel, error) {
	c, err := fc.repo.GetVids(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (fc *FetcherCase) GetVideos(c *pb.Channel) ([]*pb.Video, error) {
	videos, err := fc.repo.GetVideos(c)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (fc *FetcherCase) GetVideo(v *pb.Video) (*pb.Video, error) {
	if v.Vid == "" {
		return nil, errors.New("fc.GetVideo err: video id is nil")
	}
	video, err := fc.repo.GetVideo(v)
	if err != nil {
		return nil, err
	}
	return video, nil
}
