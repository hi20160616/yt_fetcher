package mysql

import (
	"database/sql"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/pkg/errors"
)

func SelectThumbnailsByVid(db *sql.DB, vid string) ([]*pb.Thumbnail, error) {
	query := `SELECT * FROM thumbnails where vid=?`
	rows, err := db.Query(query, vid)
	if err != nil {
		return nil, errors.WithMessage(err, "SelectItemsByVid Query error")
	}
	defer rows.Close()

	ths := []*pb.Thumbnail{}
	for rows.Next() {
		var th = &pb.Thumbnail{}
		if err := rows.Scan(&th.Id, &th.Width, &th.Height, &th.URL, &th.Vid); err != nil {
			return nil, errors.WithMessage(err, "SelectItemsByVid Scan error")
		}
		ths = append(ths, th)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ths, nil
}

func InsertThumbnail(db *sql.DB, th *pb.Thumbnail) error {
	query := `INSERT INTO thumbnails(id, width, height, url, vid) VALUES(?, ?, ?, ?, ?)`
	_, err := db.Query(query, th.Id, th.Width, th.Height, th.URL, th.Vid)
	if err != nil {
		return errors.WithMessage(err, "InsertThumbnail error")
	}
	return nil
}

func UpdateThumbnail(db *sql.DB, th *pb.Thumbnail) error {
	query := "UPDATE thumbnails SET id=?, width=?, height=?, url=?, vid=? WHERE id=?;"
	if _, err := db.Exec(query, th.Id, th.Width, th.Height, th.URL, th.Vid, th.Id); err != nil {
		return err
	}
	return nil
}

func InsertOrUpdateThumbnail(db *sql.DB, th *pb.Thumbnail) error {
	exist, err := TidExist(db, th.Id)
	if err != nil {
		return err
	}
	if exist {
		return UpdateThumbnail(db, th)
	} else {
		return InsertThumbnail(db, th)
	}
}

func TidExist(db *sql.DB, tid string) (bool, error) {
	rows, err := db.Query("SELECT * FROM thumbnails WHERE id=?", tid)
	if err != nil {
		return false, errors.WithMessage(err, "tidExist Query error")
	}
	defer rows.Close()
	return rows.Next(), nil
}

func delExist(db *sql.DB, tid string) error {
	_, err := db.Exec("DELETE FROM thumbnails WHERE id=?", tid)
	if err != nil {
		return errors.WithMessage(err, "delExist thumbnails error")
	}
	return nil
}
