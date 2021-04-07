package data

import (
	"fmt"
	"log"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	db "github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
)

// go test -test.run=^TestGetVideo$
func TestGetVideo(t *testing.T) {
	fr := NewFetcherRepo()
	v := &pb.Video{Id: "zZM9YrGdiyQ"}
	v, err := fr.GetVideo(v)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(v)
}

func TestGetVideoPrivate(t *testing.T) {
	vid := "zZM9YrGdiyQ"
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
	c := &pb.Channel{Id: "UCYPvAwZP8pZhSMW8qs7cVCw"}

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
	c := &pb.Channel{Id: "UCYPvAwZP8pZhSMW8qs7cVCw"}

	got, err := fr.GetVideos(c, false)
	if err != nil {
		t.Fatal(err)
	}
	for i, video := range got.Videos {
		fmt.Println(i, ":", video.Title)
	}
}

func TestGetVideosFromTo(t *testing.T) {
	fr := NewFetcherRepo()
	vs, err := fr.GetVideosFromTo(&pb.Videos{})
	if err != nil {
		t.Error(err)
	}
	for _, v := range vs.Videos {
		fmt.Println(v.Cname, "\t", v.Title)
	}
}

func TestSearchVideos(t *testing.T) {
	fr := NewFetcherRepo()
	vs := &pb.Videos{Keywords: []string{"english"}}
	vs, err := fr.SearchVideos(vs)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(len(vs.Videos))
	for _, v := range vs.Videos {
		fmt.Println(v.Title)
	}
}
