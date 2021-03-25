package data

import (
	"fmt"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
)

func TestGetCid(t *testing.T) {
	ts := "pXV12sqXyKY"
	dc, err := db.NewDBCase()
	if err != nil {
		t.Error(err)
	}
	got, err := getCid(dc, ts, false)
	if err != nil {
		t.Error(err)
	}
	want := "UCPDis9pjXuqyI7RYLJ-TTSA"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetSetChannel(t *testing.T) {
	fr := NewFetcherRepo()
	c := &pb.Channel{Id: "UCMUnInmOkrWN4gof9KlhNmQ"}

	c, err := fr.GetChannel(c)
	if err != nil {
		t.Error(err)
	}

	if c.Name != "老高與小茉 Mr & Mrs Gao" {
		t.Errorf("got: %v", c.Name)
	}
}

func TestGetChannels(t *testing.T) {
	fr := NewFetcherRepo()
	cs := &pb.Channels{}

	if got, err := fr.GetChannels(cs); err != nil {
		t.Error(err)
	} else {
		for _, c := range got.Channels {
			fmt.Println(c)
		}
	}
}

func TestUpdateChannelFromSource(t *testing.T) {
	c := &pb.Channel{Id: "UC_IEcnNeHc_bwd92Ber-lew"}
	dc, err := db.NewDBCase()
	if err != nil {
		t.Error(err)
	}

	if err := updateChannelFromSource(dc, c, false); err != nil {
		t.Error(err)
	}
}

func TestUpdateChannels(t *testing.T) {
	fr := NewFetcherRepo()
	if err := fr.UpdateChannels(&pb.Channels{}, true); err != nil {
		t.Error(err)
	}
}
