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

func SelectVideo(db *sql.DB, vid string) (*pb.Video, error) {
	var title, description, cid, cname, last_updated sql.NullString
	err := db.QueryRow("select * from videos where vid=?", vid).Scan(&vid, &title, &description, &cid, &cname, &last_updated)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.WithMessagef(err, "no video with id %s", vid)
	case err != nil:
		return nil, err
	default:
		return &pb.Video{
			Vid:         vid,
			Title:       title.String,
			Description: description.String,
			Cid:         cid.String,
			Cname:       cname.String,
			LastUpdated: last_updated.String,
		}, nil
	}
}

func SelectVid(db *sql.DB, cid string) ([]string, error) {
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vids, nil
}

func SelectCid(db *sql.DB, vid string) (string, error) {
	stmt, err := db.Prepare("SELECT cid FROM videos WHERE vid = ?")
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
	stmtIns, err := db.Prepare("INSERT INTO videos(vid, cid) VALUES(?, ?)")
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

// Update update all fields to db except vid
func Update(db *sql.DB, v *pb.Video) error {
	if v.Vid == "" {
		return errors.New("provide nil vid")
	}
	stmt, err := db.Prepare("UPDATE videos SET title=?, description=?, cid=?, cname=?, last_updated=? WHERE vid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(v.Title, v.Description, v.Cid, v.Cname, v.LastUpdated, v.Vid); err != nil {
		return err
	}
	return nil
}

// Insert insert video to db
// Notice: No vid exist judgement here
func Insert(db *sql.DB, v *pb.Video) error {
	stmtIns, err := db.Prepare("insert into videos(vid, title, description, cid, cname, last_updated) values(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	if _, err = stmtIns.Exec(v.Vid, v.Title, v.Description, v.Cid, v.Cname, v.LastUpdated); err != nil {
		return err
	}
	return nil
}

func vidExist(db *sql.DB, vid string) (bool, error) {
	rows, err := db.Query("SELECT * FROM videos WHERE vid=?", vid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

// InsertOrUpdate determine vid exist first, if exist, update or else insert v to db.
func InsertOrUpdate(db *sql.DB, v *pb.Video) error {
	if v.Vid == "" {
		return errors.New("provide nil vid")
	}

	exist, err := vidExist(db, v.Vid)
	if err != nil {
		return err
	}
	if exist {
		return Update(db, v)
	} else {
		return Insert(db, v)
	}
}

func GetCname(db *sql.DB, c *pb.Channel) (*pb.Channel, error) {
	if c.Cid == "" {
		return nil, errors.New("mysql:GetCname: Query Cname from database with nil Cid")
	}
	var cname sql.NullString
	err := db.QueryRow("select cname from videos where cid=?", c.Cid).Scan(&cname)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.WithMessagef(err, "no video with cid %s", c.Cid)
	case err != nil:
		return nil, err
	default:
		c.Cname = cname.String
		return c, nil
	}
}

func GetChannel(db *sql.DB, c *pb.Channel) (*pb.Channel, error) {
	if c.Cid == "" {
		return nil, errors.New("mysql:GetCname: Query Cname from database with nil Cid")
	}
	var err error
	if c, err = GetCname(db, c); err != nil {
		return nil, err
	}
	if c.Vids, err = SelectVid(db, c.Cid); err != nil {
		return nil, err
	}
	return c, nil
}
