package job

import (
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	"github.com/hi20160616/yt_fetcher/internal/data"
	"github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
)

func AddOrUpdateChannel(id string) error {
	dc, err := mysql.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	c := &pb.Channel{Id: id}
	// get info from source
	c, err = data.GetChannelFromSource(c)
	// storage
	return mysql.InsertOrUpdateChannel(dc, c)
}

func DelChannel(id string) error {
	fr := data.NewFetcherRepo()
	return fr.DelChannel(&pb.Channel{Id: id})
}

func UpdateChannels(greedy bool) error {
	fr := data.NewFetcherRepo()
	fc := biz.NewFetcherCase(fr)

	// 1. get cids from database
	cs := &pb.Channels{}
	cs, err := fc.GetChannels(cs)
	if err != nil {
		return err
	}
	// 2. for range cids, get vids from video pages where cid is
	fc.SetGreedy(greedy)
	return fc.UpdateChannels(cs, fc.GetGreedy())
}
