package mysql

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func TestSearchVideos(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	// vs, err := SearchVideos(db, &pb.Videos{Keywords: []string{"english"}, Limit: 30, Before: 1617018775070789})
	vs, err := QueryVideos(db, &pb.Videos{Keywords: []string{"english"}, Limit: 30, Before: 1617011386456731})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(vs.Videos))
	for _, v := range vs.Videos {
		fmt.Printf("%d\t", v.LastUpdated)
		fmt.Println(v.Title)
	}
}

func TestQueryVideosFromTo(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	now := time.Now()
	after := now.AddDate(0, 0, -7).UnixNano() / 1000
	before := now.UnixNano() / 1000
	vs, err := QueryVideos(db, &pb.Videos{After: after, Before: before})
	if err != nil {
		t.Error(err)
	}
	flag := false
	for _, v := range vs.Videos {
		if v.Title == "Delhi Covid Crisis: Over 1000 New Cases In Last 24 Hours, Active Cases At 4890" {
			fmt.Println(v.Thumbnails)
			flag = true
		}
	}
	if !flag {
		t.Errorf("want: true, got false")
	}

}

func TestSelectVideo(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	got := &pb.Video{Id: "zyy6okYjAfE"}
	// got := &pb.Video{Id: "zZM9YrGdiyQ"}
	if got, err = SelectVideoByVid(db, got); err != nil {
		t.Errorf("err: %+v", err)
	}
	want := &pb.Video{
		Id:          "5TW7ALXdlw8",
		Title:       "專給最勇敢警探的10道神秘謎題2",
		Description: "test for description 2",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		Cname:       "亮生活 / Bright Side",
		LastUpdated: 1612601612245194,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: %v, want: %v", got, want)
	}

}

func TestInsertVideo(t *testing.T) {
	video := &pb.Video{
		Id:          "5TW7ALXdlw8",
		Title:       "專給最勇敢警探的10道神秘謎題",
		Description: "test for description",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		LastUpdated: 1612601612245194,
	}
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err := InsertVideo(db, video); err != nil {
		t.Errorf("err: %+v", err)
	}

	// test
	v := &pb.Video{Id: "5TW7ALXdlw8"}
	if got, err := SelectVideoByVid(db, v); err != nil {
		t.Error(err)
	} else {
		if video.Id == got.Id &&
			video.Title == got.Title &&
			video.Description == got.Description &&
			video.Cid == got.Cid &&
			video.LastUpdated == got.LastUpdated {
			t.Errorf("want: %+v, got: %+v", video, got)
		}
	}
}

// test via review the output
func TestSelectVidsByCid(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if vs, err := SelectVidsByCid(db, "UCCtTgzGzQSWVzCG0xR7U-MQ"); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		for _, v := range vs {
			fmt.Println(v)
		}
	}
}

func TestGetNextSearch(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	page := &Page{keywords: []string{"english"}, limit: 10}

	for {
		page, err = getNextSearch(db, page)
		if err != nil {
			t.Error(err)
		}
		for _, v := range page.videos.Videos {
			fmt.Println(v.Title)
		}
		time.Sleep(5 * time.Second)
	}
}

func TestNextSearch(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	for {
		vs, err := NextSearch(db, "english")
		if err != nil {
			t.Error(err)
		}
		for _, v := range vs.Videos {
			fmt.Println(v.Title)
		}
		time.Sleep(5 * time.Second)
	}
}
