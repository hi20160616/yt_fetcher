package mysql

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func DB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "yt_fetcher"
	cfg.Passwd = "ytpassword"
	cfg.DBName = "yt_fetcher"
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertVideo() error {
	db, err := DB()
	if err != nil {
		return err
	}

	video := &pb.Video{
		Vid:         "5TW7ALXdlw8",
		Title:       "專給最勇敢警探的10道神秘謎題",
		Cid:         "UCCtTgzGzQSWVzCG0xR7U-MQ",
		LastUpdated: "1612601612245194",
	}

	stmtIns, err := db.Prepare("insert into videos(vid, title, cid, last_updated) values(?,?,?,?)")
	if err != nil {
		return err
	}
	defer db.Close()
	defer stmtIns.Close()
	if _, err = stmtIns.Exec(video.Vid, video.Title, video.Cid, video.LastUpdated); err != nil {
		return err
	}
	return nil
}

func QVideoByVid(id string) ([]string, error) {
	db, err := DB()
	if err != nil {
		return nil, err
	}
	video := []string{}
	var title, description, cid string
	var last_updated string
	err = db.QueryRow("select * from videos where id=?", id).Scan(&id, &title, &description, &cid, &last_updated)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no video with id " + id)
	case err != nil:
		return nil, err
	default:
		video = append(video, id, title, description, cid, last_updated)
	}
	return video, nil
}

func QVidsByCid(cid string) ([]string, error) {
	db, err := DB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select id from videos where cid=?", cid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vids := make([]string, 0)

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer
			return nil, err
		}
		vids = append(vids, id)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vids, nil
}
