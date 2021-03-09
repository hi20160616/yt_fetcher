package data

import (
	"fmt"
	"log"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
)

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

func TestGetVideoFromApi(t *testing.T) {
	vid := "FE15vkiXuwE"
	want := "四千億隻蝗蟲哪去了？這可能就是人類的結局 | 老高與小茉 Mr & Mrs Gao"
	dc, err := db.NewDBCase()
	if err != nil {
		t.Error(err)
	}
	got, err := getVideoFromApi(dc, vid)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(got.Title)
	if got.Title != want {
		t.Errorf("got %v, want %v", got.Title, want)
	}
}

// go test -test.run=^TestGetVideo$
func TestGetVideo(t *testing.T) {
	fr := NewFetcherRepo()
	v := &pb.Video{Id: "FE15vkiXuwE"}
	v, err := fr.GetVideo(v)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(v)
}

func TestGetVideoPrivate(t *testing.T) {
	vid := "NHgXDqU-ihM"
	dc, err := db.NewDBCase()
	if err != nil {
		t.Error(err)
	}
	v, err := getVideo(dc, &pb.Video{Id: vid}, false)
	if err != nil {
		t.Error(err)
	}
	log.Println(v)
}

// go test -test.run=^TestGetVideoIds$
func TestGetVids(t *testing.T) {
	fr := NewFetcherRepo()
	c := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ"}

	c, err := fr.GetVids(c, false)
	if err != nil {
		t.Fatal(err)
	}

	for _, id := range c.Vids {
		fmt.Println(id)
	}

}

// go test -test.run=^TestGetVideos$
func TestGetVideos(t *testing.T) {
	fr := NewFetcherRepo()
	c := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ"}

	got, err := fr.GetVideos(c, false)
	if err != nil {
		t.Fatal(err)
	}
	for i, video := range got.Videos {
		fmt.Println(i, ":", video.Title)
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

func TestUpdateChannels(t *testing.T) {
	c := &pb.Channel{Id: "UCMUnInmOkrWN4gof9KlhNmQ"}
	dc, err := db.NewDBCase()
	if err != nil {
		t.Error(err)
	}

	if err := updateChannelFromSource(dc, c, false); err != nil {
		t.Error(err)
	}
}

func TestGetVideosFromTo(t *testing.T) {
	fr := NewFetcherRepo()
	vs := &pb.Videos{}
	vs, err := fr.GetVideosFromTo(vs)
	if err != nil {
		t.Error(err)
	}
	for _, v := range vs.Videos {
		fmt.Println(v.Cname, "\t", v.Title)
	}
}
