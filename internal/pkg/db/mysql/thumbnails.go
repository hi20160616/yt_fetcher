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
	query := `INSERT INTO thumbnails(id, width, height, url, vid) VALUES(?, ?, ?, ?, ?)` +
		` ON DUPLICATE KEY UPDATE id=?, width=?, height=?, url=?, vid=?`
	_, err := db.Exec(query,
		th.Id, th.Width, th.Height, th.URL, th.Vid,
		th.Id, th.Width, th.Height, th.URL, th.Vid)
	if err != nil {
		return errors.WithMessage(err, "InsertThumbnail error")
	}
	return nil
}

func InsertThumbnails(db *sql.DB, ths []*pb.Thumbnail) error {
	for _, th := range ths {
		err := InsertThumbnail(db, th)
		if err != nil {
			// if strings.Contains(err.Error(), "Duplicate entry") {
			//         return UpdateThumbnail(db, th)
			// }
			return err
		}
	}
	return nil
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

func VideoThumbnailsExist(db *sql.DB, vid string) (bool, error) {
	rows, err := db.Query("SELECT * FROM thumbnails WHERE vid=?", vid)
	if err != nil {
		return false, errors.WithMessage(err, "VideoThumbnailsExist Query error")
	}
	defer rows.Close()
	i := 0
	if rows.Next() {
		i += 1
	}
	if err = rows.Err(); err != nil {
		return false, err
	}
	return (i == 4), nil
}
