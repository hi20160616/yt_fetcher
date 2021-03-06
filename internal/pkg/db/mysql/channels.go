package mysql

import (
	"database/sql"
	"strconv"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/pkg/errors"
)

// SelectChannelNameByCid select name from channels by channel id
func SelectChannelByCid(db *sql.DB, c *pb.Channel) error {
	if c.Id == "" {
		return errors.New("mysql:SelectCname: Query Cname from database with nil Cid")
	}
	var name, last_updated sql.NullString
	var rank sql.NullInt32
	err := db.QueryRow("select * from channels where id=?", c.Id).Scan(&c.Id, &name, &rank, &last_updated)
	switch {
	case err == sql.ErrNoRows:
		return errors.WithMessagef(err, "SelectChannelByCid match no channel with id %v in table", c.Id)
	case err != nil:
		return errors.WithMessage(err, "SelectChannelByCid QueryRow error")
	default:
		c.Name = name.String
		c.Rank = rank.Int32
		c.LastUpdated = last_updated.String
		return nil
	}
}

func SelectChannels(db *sql.DB, cs *pb.Channels) (*pb.Channels, error) {
	rows, err := db.Query("SELECT * FROM channels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, last_updated sql.NullString
		var rank sql.NullInt32
		if err := rows.Scan(&id, &name, &rank, &last_updated); err != nil {
			return nil, err
		}
		cs.Channels = append(cs.Channels, &pb.Channel{
			Id:          id.String,
			Name:        name.String,
			Rank:        rank.Int32,
			LastUpdated: last_updated.String})
	}
	return cs, nil
}

func InsertChannel(db *sql.DB, c *pb.Channel) error {
	stmt, err := db.Prepare("INSERT INTO channels (id, name, `rank`, last_updated) values(?,?,?,?)")
	if err != nil {
		return errors.WithMessage(err, "InsertChannel stmt Prepare error")
	}
	defer stmt.Close()
	c.LastUpdated = strconv.FormatInt(time.Now().UnixNano(), 10)[:16]
	if _, err = stmt.Exec(c.Id, c.Name, c.Rank, c.LastUpdated); err != nil {
		return errors.WithMessage(err, "InsertChannel stmt.Exec error")
	}
	return nil
}

func UpdateChannel(db *sql.DB, c *pb.Channel) error {
	if c.Id == "" {
		return errors.New("provide nil id while update channel")
	}
	stmt, err := db.Prepare("UPDATE channels SET name=?, `rank`=?, last_updated=? WHERE id=?")
	if err != nil {
		return errors.WithMessage(err, "UpdateChannel stmt Prepare error")
	}
	defer stmt.Close()
	c.LastUpdated = strconv.FormatInt(time.Now().UnixNano(), 10)[:16]
	if _, err := stmt.Exec(c.Name, c.Rank, c.LastUpdated, c.Id); err != nil {
		return errors.WithMessage(err, "UpdateChannel stmt.Exec error")
	}
	return nil
}

func CidExist(db *sql.DB, cid string) (bool, error) {
	rows, err := db.Query("SELECT * FROM channels WHERE id=?", cid)
	if err != nil {
		return false, errors.WithMessage(err, "CidExist error")
	}
	defer rows.Close()
	return rows.Next(), nil
}

// InsertOrUpdateChannel determine id exist first, if exist, update or else insert c to db.
func InsertOrUpdateChannel(db *sql.DB, c *pb.Channel) error {
	if c.Id == "" {
		return errors.New("Provide nil channel id")
	}

	exist, err := CidExist(db, c.Id)
	if err != nil {
		return errors.WithMessage(err, "InsertOrUpdateChannel CidExist error")
	}
	if exist {
		return UpdateChannel(db, c)
	} else {
		return InsertChannel(db, c)
	}

}

func DelChannel(db *sql.DB, c *pb.Channel) error {
	if c.Id == "" {
		return errors.New("Provide nil channel id")
	}

	exist, err := CidExist(db, c.Id)
	if err != nil {
		return err
	}

	if exist {
		row := db.QueryRow("DELETE FROM channels WHERE id = ?", c.Id)
		if row.Err() == nil {
			row = db.QueryRow("DELETE FROM videos WHERE cid = ?", c.Id)
		}
		return row.Err()
	}
	return errors.New("DelChannel error")
}
