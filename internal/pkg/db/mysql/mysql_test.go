package mysql

import (
	"fmt"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func TestInsertVideo(t *testing.T) {
	video := &pb.Video{
		Vid:         "5TW7ALXdlw8",
		Title:       "專給最勇敢警探的10道神秘謎題",
		Description: "test for description",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		Cname:       "亮生活 / Light Side",
		LastUpdated: "1612601612245194",
	}
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err := Insert(db, video); err != nil {
		t.Errorf("err: %+v", err)
	}
}

func TestInsertOrUpdate(t *testing.T) {
	video := &pb.Video{
		Vid:         "5TW7ALXdlw8",
		Title:       "專給最勇敢警探的10道神秘謎題",
		Description: "test for description 1",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		Cname:       "亮生活 / Light Side",
		LastUpdated: "1612601612245194",
	}

	tc := struct {
		video *pb.Video
		want  string
	}{
		video, "test for description 1",
	}

	dc, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer dc.Close()
	if err := InsertOrUpdate(dc, tc.video); err != nil {
		t.Error(err)
	}
	if v, err := SelectVideo(dc, tc.video.Vid); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		if v.Description != tc.want {
			t.Errorf("want: %s, got: %s", tc.want, v.Description)
		}
	}
}

func TestSelectVideo(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if v, err := SelectVideo(db, "5TW7ALXdlw8"); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		fmt.Println(v)
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

func TestUpdateVideo(t *testing.T) {
	video := &pb.Video{
		Vid:         "5TW7ALXdlw8",
		Title:       "test title update",
		Description: "test for description",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		Cname:       "亮生活 / Light Side",
		LastUpdated: "1612601612245194",
	}
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err := Update(db, video); err != nil {
		t.Errorf("err: %+v", err)
	}
}
