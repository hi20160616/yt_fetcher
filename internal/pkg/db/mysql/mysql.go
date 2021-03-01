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

// InsertOrUpdate determine if vid or cid exist in db, else update or insert to db.
func InsertOrUpdate(db *sql.DB, v *pb.Video, c *pb.Channel) error {
	if v.Id == "" {
		return errors.New("provide nil vid")
	}

	vExist, err := vidExist(db, v.Id)
	if err != nil {
		return err
	}
	cExist, err := cidExist(db, v.Cid)
	if err != nil {
		return err
	}

	if vExist {
		err = UpdateVideo(db, v)
	} else {
		err = InsertVideo(db, v)
	}

	if cExist {
		err = UpdateChannel(db, c)
	} else {
		err = InsertChannel(db, c)
	}
	return err
}

// InsertOrUpdateVC determine if vid or cid exist in db, else update or insert to db.
// NOTICE: this funciton will insert and update videos and channels tables in db.
// NOTICE: this function should invoke after v was populated completely.
func InsertOrUpdateVC(db *sql.DB, v *pb.Video) error {
	if v.Id == "" {
		return errors.New("provide nil vid")
	}

	vExist, err := vidExist(db, v.Id)
	if err != nil {
		return err
	}
	cExist, err := cidExist(db, v.Cid)
	if err != nil {
		return err
	}

	if vExist {
		err = UpdateVideo(db, v)
	} else {
		err = InsertVideo(db, v)
	}

	if cExist {
		err = UpdateChannel(db, &pb.Channel{Id: v.Cid, Name: v.Cname})
	} else {
		err = InsertChannel(db, &pb.Channel{Id: v.Cid, Name: v.Cname})
	}
	return err
}
