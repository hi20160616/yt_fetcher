package mysql

import (
	"database/sql"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/pkg/errors"
)

// Search just search keywords is contained in title or description
// TODO: pass test
func SearchVideos(db *sql.DB, vs *pb.Videos, keywords ...string) (*pb.Videos, error) {
	// query prapare
	query := `SELECT v.id, v.title, v.thumbnails, v.description, v.duration, v.cid, c.name AS cname, v.last_updated
		FROM videos AS v LEFT JOIN channels AS c ON v.cid = c.id`
	if len(keywords) != 0 {
		query += " WHERE "
	}
	condition := "v.title LIKE ? OR v.description LIKE ?"
	args := []interface{}{}
	for i, v := range keywords {
		if i != 0 {
			query += " OR "
		}
		query += condition
		args = append(args, "%"+v+"%", "%"+v+"%")
	}

	// query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, errors.WithMessage(err, "Search error")
	}
	defer rows.Close()

	// populate
	if err = selectVideos(db, vs, rows); err != nil {
		return nil, err
	}
	return vs, nil
}

// SelectVideosFromTo select videos left join channels where rank != -1
// TODO: pass test
func SelectVideosFromTo(db *sql.DB, vs *pb.Videos) (*pb.Videos, error) {
	q := "SELECT v.id, v.title, v.description, v.duration, v.cid, c.name AS cname, v.last_updated FROM videos AS v LEFT JOIN channels AS c on v.cid = c.id WHERE v.last_updated>? AND v.last_updated<? AND c.rank<>-1 order by cid;"
	rows, err := db.Query(q, vs.After, vs.Before)
	if err != nil {
		return nil, errors.WithMessage(err, "SelectVideosFromTo error")
	}
	defer rows.Close()

	if err = selectVideos(db, vs, rows); err != nil {
		return nil, err
	}
	return vs, nil
}

func SelectVideosByCid(db *sql.DB, channelId string) (*pb.Videos, error) {
	q := "SELECT v.id, v.title, v.description, v.duration, v.cid, c.name AS cname, v.last_updated FROM videos AS v LEFT JOIN channels AS c on v.cid=c.id WHERE c.id=?;"
	rows, err := db.Query(q, channelId)
	if err != nil {
		return nil, errors.WithMessage(err, "SelectVideosByCid query error")
	}
	defer rows.Close()

	videos := &pb.Videos{}
	if err = selectVideos(db, videos, rows); err != nil {
		return nil, err
	}
	return videos, nil
}

// SelectVideoById select video from videos by video id
func SelectVideoByVid(db *sql.DB, v *pb.Video) (*pb.Video, error) {
	q := "SELECT v.id, v.title, v.description, v.duration, v.cid, c.name AS cname, v.last_updated FROM videos AS v LEFT JOIN channels AS c on v.cid=c.id WHERE v.id=?;"
	rows, err := db.Query(q, v.Id)
	if err != nil {
		return nil, err
	}
	vs := &pb.Videos{}
	if err = selectVideos(db, vs, rows); err != nil {
		return nil, err
	}
	if len(vs.Videos) == 0 {
		return nil, errors.New("SelectVideoByVid got nil from id: " + v.Id)
	}
	return vs.Videos[0], nil
}

func selectVideos(db *sql.DB, videos *pb.Videos, rows *sql.Rows) error {
	var id, title, description, duration, cid, cname, last_updated sql.NullString
	for rows.Next() {
		if err := rows.Scan(&id, &title, &description, &duration, &cid, &cname, &last_updated); err != nil {
			return errors.WithMessage(err, "SelectVideosByCid rows.Scan error")
		}

		ths, err := SelectThumbnailsByVid(db, id.String)
		if err != nil {
			return err
		}
		videos.Videos = append(videos.Videos, &pb.Video{
			Id:          id.String,
			Title:       title.String,
			Thumbnails:  ths,
			Description: description.String,
			Duration:    duration.String,
			Cid:         cid.String,
			Cname:       cname.String,
			LastUpdated: last_updated.String,
		})
	}

	if err := rows.Err(); err != nil {
		return errors.WithMessage(err, "SelectVideosByCid rows error")
	}
	return nil
}

// SelectVidsByCid select video id list from videos by channel id
func SelectVidsByCid(db *sql.DB, cid string) ([]string, error) {
	rows, err := db.Query("select id from videos where cid=?", cid)
	if err != nil {
		return nil, errors.WithMessage(err, "SelectVidsByCid Query error")
	}
	defer rows.Close()
	vids := make([]string, 0)

	for rows.Next() {
		var vid string
		if err := rows.Scan(&vid); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer
			return nil, errors.WithMessage(err, "SelectVidsByCid Scan error")
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
		return "", errors.WithMessage(err, "SelectCidByVid stmt Prepare error")
	}
	defer stmt.Close()
	var cid string
	if err = stmt.QueryRow(vid).Scan(&cid); err != nil {
		return "", errors.WithMessage(err, "SelectCidByVid stmt QueryRow error")
	}
	return cid, nil
}

// InsertVids insert vids with cid
func InsertVids(db *sql.DB, vids []string, cid string) error {
	for _, vid := range vids {
		exist, err := VidExist(db, vid)
		if err != nil {
			return errors.WithMessage(err, "InsertVids error")
		}
		if exist {
			continue
		}
		_, err = db.Exec("INSERT INTO videos(id, cid) VALUES(?, ?)", vid, cid)
		if err != nil {
			return errors.WithMessage(err, "InsertVids Exec error")
		}
	}
	return nil
}

// UpdateVideo update all fields to db except vid
func UpdateVideo(db *sql.DB, v *pb.Video) error {
	if v.Id == "" {
		return errors.New("provide nil vid")
	}
	_, err := db.Exec("UPDATE videos SET title=?, description=?, duration=?, cid=?, last_updated=? WHERE id=?",
		v.Title, v.Description, v.Duration, v.Cid, v.LastUpdated, v.Id)
	if err != nil {
		return errors.WithMessage(err, "UpdateVideo Exec error")
	}
	return nil
}

// InsertVideo insert video to db
// Notice: No vid exist judgement here
func InsertVideo(db *sql.DB, v *pb.Video) error {
	_, err := db.Exec(
		"INSERT INTO videos(id, title, description, duration, cid, last_updated) VALUES(?,?,?,?,?,?)",
		v.Id, v.Title, v.Description, v.Duration, v.Cid, v.LastUpdated)
	if err != nil {
		return errors.WithMessage(err, "InsertVideo Exec error")
	}
	return nil
}

func VidExist(db *sql.DB, vid string) (bool, error) {
	rows, err := db.Query("SELECT * FROM videos WHERE id=?", vid)
	if err != nil {
		return false, errors.WithMessage(err, "VidExist Query error")
	}
	defer rows.Close()
	return rows.Next(), nil
}

// InsertOrUpdateVideo determine vid exist first, if exist, update or else insert v to db.
func InsertOrUpdateVideo(db *sql.DB, v *pb.Video) error {
	if v.Id == "" {
		return errors.New("provide nil videoId")
	}

	exist, err := VidExist(db, v.Id)
	if err != nil {
		return errors.WithMessage(err, "InsertOrUpdateVideo VidExist error")
	}
	if exist {
		return UpdateVideo(db, v)
	} else {
		return InsertVideo(db, v)
	}
}
