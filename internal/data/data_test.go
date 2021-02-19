package data

import (
	"fmt"
	"log"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
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
	got, err := getCid(ts)
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
	v := &pb.Video{Vid: "-2u6RirE7aI"}
	v, err := fr.GetVideo(v)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(v)
}

// go test -test.run=^TestGetVideoIds$
func TestGetVids(t *testing.T) {
	fr := NewFetcherRepo()
	c := &pb.Channel{Cid: "UCCtTgzGzQSWVzCG0xR7U-MQ"}

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
	c := &pb.Channel{Cid: "UCCtTgzGzQSWVzCG0xR7U-MQ"}

	_, err := fr.GetVideos(c)
	if err != nil {
		t.Fatal(err)
	}
	// log.Println(v)
}
