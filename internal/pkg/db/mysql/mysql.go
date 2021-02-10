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

func InsertVids(vids []string, cid string) error {
	db, err := DB()
	if err != nil {
		return err
	}

	stmtIns, err := db.Prepare("INSERT INTO videos(vid, cid) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer db.Close()
	defer stmtIns.Close()
	for _, vid := range vids {
		if _, err = stmtIns.Exec(vid, cid); err != nil {
			return err
		}
	}
	return err
}

// TODO: pass test
func UpdateVideo(v *pb.Video) error {
	db, err := DB()
	if err != nil {
		return nil
	}

	stmt, err := db.Prepare("UPDATE videos SET title=?, description=?, cid=?, cname=?, last_updated=? WHERE vid=?")
	if err != nil {
		return err
	}
	defer db.Close()
	defer stmt.Close()
	if _, err := stmt.Exec(v.Title, v.Description, v.Cid, v.Cname, v.LastUpdated, v.Vid); err != nil {
		return err
	}
	return nil
}

// TODO: how to deal duplicated vid
// TODO: rename to AddVideo
func InsertVideo(v *pb.Video) error {
	db, err := DB()
	if err != nil {
		return err
	}

	// TODO: if vid exist, update the values, or not exist, insert

	stmtIns, err := db.Prepare("insert into videos(vid, title, description, cid, cname, last_updated) values(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer db.Close()
	defer stmtIns.Close()
	if _, err = stmtIns.Exec(v.Vid, v.Title, v.Description, v.Cid, v.Cname, v.LastUpdated); err != nil {
		return err
	}
	return nil
}

func VidExist(vid string) (bool, error) {
	db, err := DB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM videos WHERE vid=?", vid)
	if err != nil {
		return false, err
	}
	return rows.Next(), nil
}

func QVideoByVid(vid string) ([]string, error) {
	db, err := DB()
	if err != nil {
		return nil, err
	}
	video := []string{}
	var title, description, cid, cname, last_updated sql.NullString
	err = db.QueryRow("select * from videos where vid=?", vid).Scan(&vid, &title, &description, &cid, &cname, &last_updated)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no video with id " + vid)
	case err != nil:
		return nil, err
	default:
		video = append(video, vid, title.String, description.String, cid.String, cname.String, last_updated.String)
	}
	return video, nil
}

func QVidsByCid(cid string) ([]string, error) {
	db, err := DB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select vid from videos where cid=?", cid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vids := make([]string, 0)

	for rows.Next() {
		var vid string
		if err := rows.Scan(&vid); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer
			return nil, err
		}
		vids = append(vids, vid)
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
