package mysql

import (
	"fmt"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func TestInsertThumbnail(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	th := &pb.Thumbnail{
		Id:     "FE15vkiXuwE_w168",
		Width:  168,
		Height: 94,
		URL:    "https://i.ytimg.com/vi/FE15vkiXuwE/hqdefault.jpg?sqp=-oaymwEiCKgBEF5IWvKriqkDFQgBFQAAAAAYASUAAMhCPQCAokN4AQ==&rs=AOn4CLCpHEZcnlM2PdSwf3OG1A0FLb1a7w",
		Vid:    "FE15vkiXuwE",
	}

	if err = delExist(db, th.Id); err != nil {
		t.Error(err)
	}

	if err = InsertThumbnail(db, th); err != nil {
		t.Error(err)
	}

	gots, err := SelectThumbnailsByVid(db, th.Vid)
	if err != nil {
		t.Error(err)
	}
	for _, got := range gots {
		if got.Id == th.Id {
			if got.Width == th.Width &&
				got.Height == th.Height &&
				got.URL == th.URL &&
				got.Vid == th.Vid {
				fmt.Println("pass test")
			} else {
				t.Errorf("got: %v", got)
			}
		}
	}
}

func TestInsertOrUpdateThumbnail(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	th := &pb.Thumbnail{
		Id:     "FE15vkiXuwE_w168",
		Width:  168,
		Height: 94,
		URL:    "https://i.ytimg.com/vi/FE15vkiXuwE/hqdefault.jpg?sqp=-oaymwEiCKgBEF5IWvKriqkDFQgBFQAAAAAYASUAAMhCPQCAokN4AQ==&rs=AOn4CLCpHEZcnlM2PdSwf3OG1A0FLb1a7w",
		Vid:    "FE15vkiXuwE",
	}

	if err = InsertOrUpdateThumbnail(db, th); err != nil {
		t.Error(err)
	}

	gots, err := SelectThumbnailsByVid(db, th.Vid)
	if err != nil {
		t.Error(err)
	}
	for _, got := range gots {
		if got.Id == th.Id {
			if got.Width == th.Width &&
				got.Height == th.Height &&
				got.URL == th.URL &&
				got.Vid == th.Vid {
				fmt.Println("pass test")
			} else {
				t.Errorf("got: %v", got)
			}
		}
	}

}

func TestTidExist(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	th := &pb.Thumbnail{
		Id:     "FE15vkiXuwE_w168",
		Width:  168,
		Height: 944,
		URL:    "https://i.ytimg.com/vi/FE15vkiXuwE/hqdefault.jpg?sqp=-oaymwEiCKgBEF5IWvKriqkDFQgBFQAAAAAYASUAAMhCPQCAokN4AQ==&rs=AOn4CLCpHEZcnlM2PdSwf3OG1A0FLb1a7w",
		Vid:    "FE15vkiXuwE",
	}
	exist, err := TidExist(db, th.Id)
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("got %v", exist)
	}
}
