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
	if err := InsertVideo(video); err != nil {
		t.Errorf("err: %+v", err)
	}
}

func TestQVideoById(t *testing.T) {
	if v, err := QVideoByVid("5TW7ALXdlw8"); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		fmt.Println(v)
	}

}

func TestQVidsByChannelId(t *testing.T) {
	if vs, err := QVidsByCid("UCCtTgzGzQSWVzCG0xR7U-MQ"); err != nil {
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
	if err := UpdateVideo(video); err != nil {
		t.Errorf("err: %+v", err)
	}
}
