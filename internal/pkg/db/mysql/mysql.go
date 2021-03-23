package mysql

import (
	"database/sql"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/pkg/errors"
)

func NewDBCase() (*sql.DB, error) {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	o, err := exec.Command("enit", "get", "yt_fetcher").Output()
	if err != nil {
		return nil, err
	}
	c := strings.TrimSpace(strings.Split(string(o), "=")[1])
	if c == "" {
		return nil, errors.New("SQL conn string is nil")
	}

	db, err := sql.Open("mysql", string(c))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InsertOrUpdate insert or update video and thumbnails by v and channel by c
func InsertOrUpdate(db *sql.DB, v *pb.Video, c *pb.Channel) error {
	if v.Id == "" {
		return errors.New("provide nil vid")
	}

	if err := InsertOrUpdateThumbnails(db, v.Thumbnails); err != nil {
		return err
	}
	if err := InsertOrUpdateVideo(db, v); err != nil {
		return err
	}
	return InsertOrUpdateChannel(db, c)
}

// InsertOrUpdate2 insert or update videos, thumbnails, channels by v
// NOTICE: this function should invoke after v was populated completely.
func InsertOrUpdate2(db *sql.DB, v *pb.Video) error {
	return InsertOrUpdate(db, v, &pb.Channel{Id: v.Cid, Name: v.Cname})
}
