package data

import (
	"fmt"
	"strings"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
)

func TestGetChannelFromSource(t *testing.T) {
	c := &pb.Channel{Id: "UCMUnInmOkrWN4gof9KlhNmQ"}

	c, err := getChannelFromSource(c)
	if err != nil {
		t.Error(err)
	}

	if c.Name != "Mr & Mrs Gao" {
		t.Errorf("got: %v", c.Name)
	}
}

func TestGetVideoFromApi(t *testing.T) {
	vid := "ejXYnJxNKOI"
	want := "四千億隻蝗蟲哪去了？這可能就是人類的結局 | 老高與小茉 Mr & Mrs Gao"
	dc, err := db.NewDBCase()
	if err != nil {
		t.Error(err)
	}
	got, err := getVideoFromApi(dc, vid)
	if err != nil {
		if strings.Contains(err.Error(), "LIVE_STREAM_OFFLINE") {
			fmt.Println("pass")
			return
		}
	}
	fmt.Println(got.Title)
	if got.Title != want {
		t.Errorf("got %v, want %v", got.Title, want)
	}
	for _, v := range got.Thumbnails {
		fmt.Println(v)
	}
}

func TestGetCidFromSource(t *testing.T) {
	tvid := "pXV12sqXyKY"
	got, err := getCidFromSource(tvid)
	if err != nil {
		t.Error(err)
	}
	want := "UCPDis9pjXuqyI7RYLJ-TTSA"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
