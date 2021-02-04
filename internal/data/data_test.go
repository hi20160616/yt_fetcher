package data

import (
	"log"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func TestGetVideo(t *testing.T) {
	f := NewFetcherRepo()
	v := &pb.Video{Id: "nyfAij5B9fM"}
	v, err := f.GetVideo(v)
	if err != nil {
		t.Fatal(err)
	}
	// log.Println(v.GetTitle())
	log.Println(v)
}
