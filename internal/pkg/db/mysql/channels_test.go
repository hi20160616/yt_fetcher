package mysql

import (
	"fmt"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func TestInsertChannel(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	c := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ", Name: "亮生活 / Bright Side"}
	if err = InsertChannel(db, c); err != nil {
		t.Errorf("err: %+v", err)
	}
	cc := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ"}
	if err := SelectChannelByCid(db, cc); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		if cc.Name != "亮生活 / Bright Side" {
			t.Errorf("got: %s", c.Name)
		}
	}
}

func TestSelectChannelName(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	c := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ"}
	if err := SelectChannelByCid(db, c); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		if c.Name != "亮生活 / Bright Side" {
			t.Errorf("got: %s", c.Name)
		}
	}
}

func TestSelectChannels(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	cs := &pb.Channels{}
	got, err := SelectChannels(db, cs)
	for _, c := range got.Channels {
		fmt.Println(c)
	}
}

func TestInsertOrUpdateChannel(t *testing.T) {
	db, err := NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	c := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ", Name: "亮生活 / Bright Side"}
	if err = InsertOrUpdateChannel(db, c); err != nil {
		t.Errorf("err: %+v", err)
	}
	cc := &pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ"}
	if err := SelectChannelByCid(db, cc); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		if cc.Name != "亮生活 / Bright Side" {
			t.Errorf("got: %s", c.Name)
		}
	}

}
