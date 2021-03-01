package mysql

import (
	"database/sql"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/pkg/errors"
)

func SelectVideosByCid(db *sql.DB, channelId string) ([]*pb.Video, error) {
	var id, title, description, cid, cname, last_updated sql.NullString

	rows, err := db.Query("SELECT v.id, v.title, v.description, v.cid, c.name AS cname, v.last_updated FROM videos AS v LEFT JOIN channels AS c on v.cid=c.id WHERE c.id=?;", channelId)
	if err != nil {
		return nil, errors.WithMessage(err, "SelectVideosByCid query error")
	}
	defer rows.Close()

	videos := make([]*pb.Video, 0)
	for rows.Next() {
		if err := rows.Scan(&id, &title, &description, &cid, &cname, &last_updated); err != nil {
			return nil, errors.WithMessage(err, "SelectVideosByCid rows.Scan error")
		}
		videos = append(videos, &pb.Video{Id: id.String,
			Title:       title.String,
			Description: description.String,
			Cid:         cid.String,
			Cname:       cname.String,
			LastUpdated: last_updated.String,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "SelectVideosByCid rows error")
	}
	return videos, nil
}

// SelectVideoById select video from videos by video id
func SelectVideoByVid(db *sql.DB, v *pb.Video) error {
	var title, description, cid, cname, last_updated sql.NullString
	// err := db.QueryRow("select * from videos where id=?", v.Id).Scan(&v.Id, &title, &description, &cid, &last_updated)
	err := db.QueryRow("SELECT v.id, v.title, v.description, v.cid, c.name AS cname, v.last_updated FROM videos AS v LEFT JOIN channels AS c on v.cid=c.id WHERE v.id=?;", v.Id).Scan(&v.Id, &title, &description, &cid, &cname, &last_updated)
	switch {
	case err == sql.ErrNoRows:
		return errors.WithMessagef(err, "no video with id %s", v.Id)
	case err != nil:
		return errors.WithMessage(err, "SelectVideoByVid QueryRow error")
	default:
		v.Title = title.String
		v.Description = description.String
		v.Cid = cid.String
		v.Cname = cname.String
		v.LastUpdated = last_updated.String
		return nil
	}
}

// SelectVidsByCid select video id list from videos by channel id
func SelectVidsByCid(db *sql.DB, cid string) ([]string, error) {
	rows, err := db.Query("select id from videos where cid=?", cid)
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vids, nil
}

// SelectCidByVid select channel id from videos by video id
func SelectCidByVid(db *sql.DB, vid string) (string, error) {
	stmt, err := db.Prepare("SELECT cid FROM videos WHERE id = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var cid string
	if err = stmt.QueryRow(vid).Scan(&cid); err != nil {
		return "", err
	}
	return cid, nil
}

// InsertVids insert vids with cid
func InsertVids(db *sql.DB, vids []string, cid string) error {
	stmtIns, err := db.Prepare("INSERT INTO videos(id, cid) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	for _, vid := range vids {
		exist, err := vidExist(db, vid)
		if err != nil {
			return err
		}
		if exist {
			continue
		}
		if _, err = stmtIns.Exec(vid, cid); err != nil {
			return err
		}
	}
	return err
}

// UpdateVideo update all fields to db except vid
func UpdateVideo(db *sql.DB, v *pb.Video) error {
	if v.Id == "" {
		return errors.New("provide nil vid")
	}
	stmt, err := db.Prepare("UPDATE videos SET title=?, description=?, cid=?, last_updated=? WHERE id=?")
	if err != nil {
		return errors.WithMessage(err, "UpdateVideo stmt Prepare error")
	}
	defer stmt.Close()
	if _, err := stmt.Exec(v.Title, v.Description, v.Cid, v.LastUpdated, v.Id); err != nil {
		return errors.WithMessage(err, "UpdateVideo stmt Exec error")
	}
	return nil
}

// InsertVideo insert video to db
// Notice: No vid exist judgement here
func InsertVideo(db *sql.DB, v *pb.Video) error {
	stmt, err := db.Prepare("insert into videos(id, title, description, cid, last_updated) values(?,?,?,?,?)")
	if err != nil {
		return errors.WithMessage(err, "InsertVideo stmt Prepare error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(v.Id, v.Title, v.Description, v.Cid, v.LastUpdated); err != nil {
		return errors.WithMessage(err, "InsertVideo stmt Exec error")
	}
	return nil
}

func vidExist(db *sql.DB, vid string) (bool, error) {
	rows, err := db.Query("SELECT * FROM videos WHERE id=?", vid)
	if err != nil {
		return false, errors.WithMessage(err, "vidExist Query error")
	}
	defer rows.Close()
	return rows.Next(), nil
}

// InsertOrUpdateVideo determine vid exist first, if exist, update or else insert v to db.
func InsertOrUpdateVideo(db *sql.DB, v *pb.Video) error {
	if v.Id == "" {
		return errors.New("provide nil videoId")
	}

	exist, err := vidExist(db, v.Id)
	if err != nil {
		return errors.WithMessage(err, "InsertOrUpdateVideo vidExist error")
	}
	if exist {
		return UpdateVideo(db, v)
	} else {
		return InsertVideo(db, v)
	}
}
