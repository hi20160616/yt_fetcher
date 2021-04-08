package mysql

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/pkg/errors"
)

type Page struct {
	keywords []string
	query    string
	args     []interface{}
	offset   int64
	limit    int
	videos   *pb.Videos
}

var P *Page = &Page{limit: 30}

// queryNextSearch search for infinite search
func queryNextSearch(db *sql.DB, page *Page) (*Page, error) {
	rows, err := db.Query(page.query, page.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := &pb.Videos{}

	if err = mkVideos(db, vs, rows); err != nil {
		return nil, err
	}

	page.videos = vs

	if len(page.videos.Videos) == page.limit {
		page.offset = page.videos.Videos[page.limit-1].LastUpdated
	}
	return page, nil
}

// getNextSearch search for infinite search
func getNextSearch(db *sql.DB, page *Page) (*Page, error) {
	// first loop
	if page.offset == 0 {
		q := `SELECT x.id, x.title, x.description, x.duration, x.cid, x.cname, x.last_updated 
			FROM (SELECT v.id, v.title, v.description, v.duration, v.cid, c.name AS cname, v.last_updated 
				FROM videos AS v LEFT JOIN channels AS c ON v.cid=c.id) AS x WHERE `
		if len(page.keywords) != 0 {
			q += " ( "
		}
		page.offset = time.Now().UnixNano() / 1000
		for i, v := range page.keywords {
			if i != 0 {
				q += " OR "
			}
			q += "x.title LIKE ? OR x.description LIKE ?"
			page.args = append(page.args, "%"+v+"%", "%"+v+"%")
		}
		if len(page.keywords) != 0 {
			q += " ) AND "
		}
		q += " x.last_updated<? ORDER BY x.last_updated DESC, v.duration DESC LIMIT ?"
		page.query = q
		page.args = append(page.args, page.offset, page.limit)
	} else {
		page.args[len(page.args)-2] = page.offset
		page.args[len(page.args)-1] = page.limit
	}
	fmt.Println(page.offset)

	// query
	return queryNextSearch(db, page)
}

// NextSearch just for infinite search
func NextSearch(db *sql.DB, keywords ...string) (*pb.Videos, error) {
	P.keywords = append(P.keywords, keywords...)
	ps, err := getNextSearch(db, P)
	if err != nil {
		return nil, err
	}
	return ps.videos, nil
}

// QueryVideos combine query by keywords and last_updated
// if vs have `after` and `before` set, query rows from `after` to `before`
// by default, `before` is now().UnixNano()/1000
// if `before` is set, it can query rows paginated to 30 rows
// if `before` and `limit` are set, it can query rows paginated by number of `limit` defined.
// if keywords are non-nil, query rows that title or description like the keywords.
func QueryVideos(db *sql.DB, vs *pb.Videos) (*pb.Videos, error) {
	args := []interface{}{}
	// query prapare
	q := `SELECT x.id, x.title, x.description, x.duration, x.cid, x.cname, x.last_updated 
			FROM (SELECT v.id, v.title, v.description, v.duration, v.cid, c.name AS cname, v.last_updated 
				FROM videos AS v LEFT JOIN channels AS c ON v.cid=c.id) AS x WHERE `
	if len(vs.Keywords) != 0 {
		q += " ( "
		for i, v := range vs.Keywords {
			if i != 0 {
				q += " OR "
			}
			q += " x.title LIKE ? OR x.description LIKE ? "
			args = append(args, "%"+v+"%", "%"+v+"%")
		}
		q += " ) AND "
	}
	q += " (x.last_updated>? AND x.last_updated<?) ORDER BY x.last_updated DESC, v.duration DESC LIMIT ? "
	// set default conditions
	if vs.Limit == -1 { // no rows limit
		q = q[:len(q)-7]
		args = append(args, vs.After, vs.Before)
	}
	if vs.Limit == 0 {
		vs.Limit = 30
	}
	if vs.Before == 0 {
		vs.Before = time.Now().UnixNano() / 1000
	}

	// query
	args = append(args, vs.After, vs.Before, vs.Limit)
	return query(db, vs, q, args)
}

func query(db *sql.DB, vs *pb.Videos, q string, args ...interface{}) (*pb.Videos, error) {
	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = mkVideos(db, vs, rows); err != nil {
		return nil, err
	}

	if len(vs.Videos) == int(vs.Limit) {
		vs.Before = vs.Videos[int(vs.Limit)-1].LastUpdated
	}
	return vs, nil
}

// SelectVideoById select video from videos by video id
func SelectVideoByVid(db *sql.DB, v *pb.Video) (*pb.Video, error) {
	q := "SELECT v.id, v.title, v.description, v.duration, v.cid, c.name AS cname, v.last_updated FROM videos AS v LEFT JOIN channels AS c on v.cid=c.id WHERE v.id=?"
	vs, err := query(db, &pb.Videos{}, q, v.Id)
	if err != nil {
		return nil, err
	}
	return vs.Videos[0], nil
}

func mkVideos(db *sql.DB, videos *pb.Videos, rows *sql.Rows) error {
	var id, title, description, duration, cid, cname sql.NullString
	var last_updated sql.NullInt64
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
			LastUpdated: last_updated.Int64,
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

// SelectVidsTidNull get vids that the tid is null in thumbnails
func SelectVidsTidNull(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT v.id AS vid FROM videos AS v LEFT JOIN thumbnails AS t ON v.id = t.vid WHERE t.id is NULL;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vids := make([]string, 0)
	for rows.Next() {
		var vid string
		if err = rows.Scan(&vid); err != nil {
			return nil, err
		}
		vids = append(vids, vid)
	}

	if err = rows.Err(); err != nil {
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

// InsertVideo insert video to db
// Notice: No vid exist judgement here
func InsertVideo(db *sql.DB, v *pb.Video) error {
	q := "INSERT INTO videos(id, title, description, duration, cid, last_updated) VALUES(?,?,?,?,?,?)" +
		" ON DUPLICATE KEY UPDATE id=?, title=?, description=?, duration=?, cid=?, last_updated=?"
	_, err := db.Exec(q,
		v.Id, v.Title, v.Description, v.Duration, v.Cid, v.LastUpdated,
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

func DelNilVideos(db *sql.DB) error {
	q := "delete from videos where duration = 0"
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
