package mysql

import (
	"fmt"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func TestSelectVideosByCid(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	cid := "UCMUnInmOkrWN4gof9KlhNmQ"
	vs, err := SelectVideosByCid(db, cid)
	if err != nil {
		t.Error(err)
	}
	flag := false
	for _, v := range vs {
		if v.Title == "目前最火的一期，消失的同學，不是平行宇宙也不是曼德拉效應，而是〇〇 | 老高與小茉 Mr & Mrs Gao" {
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

	v := &pb.Video{Id: "5TW7ALXdlw8"}
	if err := SelectVideoByVid(db, v); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		fmt.Println(v)
	}

}

func TestInsertVideo(t *testing.T) {
	video := &pb.Video{
		Id:          "5TW7ALXdlw8",
		Title:       "專給最勇敢警探的10道神秘謎題",
		Description: "test for description",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		LastUpdated: "1612601612245194",
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
	if err := SelectVideoByVid(db, v); err != nil {
		t.Error(err)
	} else {
		if video.Id == v.Id &&
			video.Title == v.Title &&
			video.Description == v.Description &&
			video.Cid == v.Cid &&
			video.LastUpdated == v.LastUpdated {
			t.Errorf("want: %+v, got: %+v", video, v)
		}
	}
}

func TestUpdateVideo(t *testing.T) {
	video := &pb.Video{
		Id:          "5TW7ALXdlw8",
		Title:       "test title update",
		Description: "test for description",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		LastUpdated: "1612601612245194",
	}
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err := UpdateVideo(db, video); err != nil {
		t.Errorf("err: %+v", err)
	}
}

func TestInsertOrUpdateVideo(t *testing.T) {
	video := &pb.Video{
		Id:          "5TW7ALXdlw8",
		Title:       "專給最勇敢警探的10道神秘謎題2",
		Description: "test for description 2",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		LastUpdated: "1612601612245194",
	}

	dc, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer dc.Close()
	if err := InsertOrUpdateVideo(dc, video); err != nil {
		t.Error(err)
	}

	// test
	v := &pb.Video{Id: "5TW7ALXdlw8"}
	if err := SelectVideoByVid(dc, v); err != nil {
		t.Error(err)
	} else {
		if video.Id == v.Id &&
			video.Title == v.Title &&
			video.Description == v.Description &&
			video.Cid == v.Cid &&
			video.LastUpdated == v.LastUpdated {
			t.Errorf("want: %+v, got: %+v", video, v)
		}
	}
}

func TestSelectVid(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if vs, err := SelectVid(db, "UCCtTgzGzQSWVzCG0xR7U-MQ"); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		for _, v := range vs {
			fmt.Println(v)
		}
	}
}
