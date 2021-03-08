package biz

import (
	"github.com/pkg/errors"

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
	GetChannelName(*pb.Channel) (*pb.Channel, error)
	GetChannel(*pb.Channel) (*pb.Channel, error)
	UpdateChannels(*pb.Channels) error
	DelChannel(*pb.Channel) error
	GetChannels(*pb.Channels) (*pb.Channels, error)
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
	if v.Id == "" {
		return nil, errors.New("fc.GetVideo err: video id is nil")
	}
	return fc.repo.GetVideo(v)
}

func (fc *FetcherCase) GetChannelName(c *pb.Channel) (*pb.Channel, error) {
	if c.Id == "" {
		return nil, errors.New("fc.GetChannel err: cid is nil")
	}
	return fc.repo.GetChannelName(c)
}

func (fc *FetcherCase) GetChannel(c *pb.Channel) (*pb.Channel, error) {
	if c.Id == "" {
		return nil, errors.New("fc.GetChannel err: cid is nil")
	}
	return fc.repo.GetChannel(c)
}

func (fc *FetcherCase) GetChannels(cs *pb.Channels) (*pb.Channels, error) {
	return fc.repo.GetChannels(cs)
}

func (fc *FetcherCase) UpdateChannels(cs *pb.Channels) error {
	return fc.repo.UpdateChannels(cs)
}
