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
	got, err := getCid(dc, ts)
	if err != nil {
		t.Error(err)
	}
	want := "UCPDis9pjXuqyI7RYLJ-TTSA"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
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
	vid := "-2u6RirE7aI"
	dc, err := db.NewDBCase()
	if err != nil {
		t.Error(err)
	}
	v, err := getVideo(dc, &pb.Video{Id: vid})
	if err != nil {
		t.Error(err)
	}
	log.Println(v)
}

// go test -test.run=^TestGetVideoIds$
func TestGetVids(t *testing.T) {
	fr := NewFetcherRepo()
	c := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ"}

	c, err := fr.GetVids(c)
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

	_, err := fr.GetVideos(c)
	if err != nil {
		t.Fatal(err)
	}
	// log.Println(v)
}

func TestGetSetChannel(t *testing.T) {
	fr := NewFetcherRepo()
	c := &pb.Channel{Id: "UCMUnInmOkrWN4gof9KlhNmQ"}

	if err := fr.GetSetChannel(c); err != nil {
		t.Error(err)
	}

	if c.Name != "Mr & Mrs Gao" {
		t.Errorf("got: %v", c.Name)
	}
}

func TestGetChannelFromSource(t *testing.T) {
	c := &pb.Channel{Id: "UCMUnInmOkrWN4gof9KlhNmQ"}

	if err := getChannelFromSource(c); err != nil {
		t.Error(err)
	}

	if c.Name != "Mr & Mrs Gao" {
		t.Errorf("got: %v", c.Name)
	}
}
