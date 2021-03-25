package data

import (
	"database/sql"
	"log"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
	"github.com/pkg/errors"
)

// getCid will get cid by vid from db,
// greedy: If greedy and cid not in db, obtain it from source of api
func getCid(dc *sql.DB, vid string, greedy bool) (string, error) {
	cid, err := db.SelectCidByVid(dc, vid)
	if !greedy {
		return cid, err
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return getCidFromSource(vid)
		}
		return "", err
	}
	return cid, nil
}

// GetChannel query channel info
// If nothing got from database, get video ids and channel info from source web page.
func (fr *fetcherRepo) GetChannel(c *pb.Channel) (*pb.Channel, error) {
	// Select name from channels
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	return getChannel(dc, c, false)
}

// getChannel get channel from database by id,
// it will fetch from source page if only greedy is true
func getChannel(dc *sql.DB, c *pb.Channel, greedy bool) (*pb.Channel, error) {
	if err := db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) && greedy { // if greedy and no rows found out
			// Get video ids and channel info from source
			c, err := getChannelFromSource(c)
			if err != nil {
				return nil, err
			}
			if err = db.InsertChannel(dc, c); err != nil { // storage channel info just got
				return nil, errors.WithMessage(err, "getChannel error")
			}
			return c, nil
		}
		return nil, err
	}
	return c, nil
}

func (fr *fetcherRepo) GetChannelName(c *pb.Channel) (*pb.Channel, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	return getChannelName(dc, c, false)
}

// getChannelName get channel name by id from database,
// it will fetch from source page if only greedy if true
func getChannelName(dc *sql.DB, c *pb.Channel, greedy bool) (*pb.Channel, error) {
	if err := db.SelectChannelByCid(dc, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) && greedy { // if greedy and no rows found out
			if c, err = getChannelFromSource(c); err != nil {
				return nil, err
			}
			return c, db.InsertChannel(dc, c) // storage channel info just got
		}
		return nil, err
	}
	return c, nil
}

// GetChannels get all channels from database
func (fr *fetcherRepo) GetChannels(cs *pb.Channels) (*pb.Channels, error) {
	dc, err := db.NewDBCase()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	return db.SelectChannels(dc, cs)
}

// DelChannel delete channel by id in c
func (fr *fetcherRepo) DelChannel(c *pb.Channel) error {
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()
	return db.DelChannel(dc, c)
}

// UpdateChannels default greedy false
func (fr *fetcherRepo) UpdateChannels(cs *pb.Channels, greedy bool) error {
	dc, err := db.NewDBCase()
	if err != nil {
		return err
	}
	defer dc.Close()

	cs, err = fr.GetChannels(cs)
	if err != nil {
		return err
	}
	for _, c := range cs.Channels {
		if err = updateChannelFromSource(dc, c, greedy); err != nil {
			return err
		}
	}
	return nil
}

// updateChannelFromSource update channel by source
// greedy true: get videos from api directly
// greedy false: get videos from api if only it is not exist in table videos
func updateChannelFromSource(dc *sql.DB, c *pb.Channel, greedy bool) error {
	c, err := getVids(dc, c, true)
	if err != nil {
		return err
	}

	iou := func(vExist, tExist bool, vid string) error {
		v, err := getVideoFromApi(dc, vid)
		if err != nil {
			return err
		}
		if vExist && !tExist {
			return db.InsertThumbnails(dc, v.Thumbnails)
		}
		return db.Insert2(dc, v)
	}

	for _, vid := range c.Vids {
		var vExist, tExist bool
		if greedy {
			// err = iou(true, true, vid)
			vExist, tExist = true, true
		} else {
			// continue on exist
			vExist, err = db.VidExist(dc, vid)
			if err != nil {
				return err
			}
			tExist, err = db.VideoThumbnailsExist(dc, vid)
			if err != nil {
				return err
			}
			// err = iou(vExist, tExist, vid)
		}
		if err = iou(vExist, tExist, vid); err != nil {
			if errors.Is(err, ErrIgnoreVideoOnPurpose) {
				log.Println(err)
				continue
			}
			return err
		}
	}
	return err
}
