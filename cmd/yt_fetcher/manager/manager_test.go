package main

import (
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/pkg/db/mysql"
)

func TestAddOrUpdateChannel(t *testing.T) {
	tcs := []struct {
		got  *pb.Channel
		want *pb.Channel
	}{
		{&pb.Channel{Id: "UCMUnInmOkrWN4gof9KlhNmQ"}, &pb.Channel{Name: "Mr & Mrs Gao"}},
		{&pb.Channel{Id: "UCCtTgzGzQSWVzCG0xR7U-MQ"}, &pb.Channel{Name: "亮生活 / Bright Side"}},
		{&pb.Channel{Id: "UCXPbm8RZhAd69OxPIV7msEg"}, &pb.Channel{Name: "HawkGuruHacker"}},
	}
	dc, err := mysql.NewDBCase()
	if err != nil {
		t.Error(err)
	}
	defer dc.Close()

	for _, tc := range tcs {
		if err := addOrUpdateChannel(tc.got.Id); err != nil {
			t.Error(err)
		}
		if err := mysql.SelectChannelByCid(dc, tc.got); err != nil {
			t.Error(err)
		}
		if tc.got.Name != tc.want.Name {
			t.Errorf("want %v, got %v", tc.want.Name, tc.got.Name)
		}
	}
}
