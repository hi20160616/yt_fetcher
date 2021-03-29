package job

import (
	"errors"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	"github.com/hi20160616/yt_fetcher/internal/data"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
)

func AddOrUpdateChannel(id string) error {
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	c := &pb.Channel{Id: id}
	// get info from source
	c, err = data.GetChannelFromSource(c)
	// storage
	return db.InsertChannel(dc, c)
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
	err = fc.UpdateChannels(cs)
	if err != nil {
		return err
	}
	// 3. rm nil videos judgement by duration = 0
	return fc.DelNilVideos()
}

func UpdateThumbnails() error {
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	vids, err := db.SelectVidsTidNull(dc)
	for _, vid := range vids {
		v, err := data.GetVideoFromApi(dc, vid)
		if err != nil {
			if errors.Is(err, data.ErrIgnoreVideoOnPurpose) {
				continue
			}
			return err
		}
		if len(v.Thumbnails) == 0 {
			return errors.New("cannot get thumbnails by videoId: " + vid)
		}
		// save it
		if err = db.InsertThumbnails(dc, v.Thumbnails); err != nil {
			return err
		}
	}
	return nil
}
