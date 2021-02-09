package mysql

import (
	"fmt"
	"testing"
)

func TestInsertVideo(t *testing.T) {
	if err := InsertVideo(); err != nil {
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
