package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() (*Database, error) {
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", host, dbname, user, password))
	if err != nil {
		return nil, err
	}

	return &Database{
		Conn: db,
	}, nil
}

func (db *Database) Select(ds *goqu.SelectDataset) (*sql.Rows, error) {
	sql, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}
	return db.Conn.Query(sql, args...)
}

func (db *Database) Insert(table string, cols []interface{}, values [][]interface{}) (*int, error) {
	now := time.Now()
	for i := range values {
		v := values[i]
		v = append(v, now, now)
		values[i] = v
	}

	ds := goqu.Insert(table).
		Cols(cols...).ColsAppend("created_at", "updated_at").
		Vals(values...).
		Returning("id")

	sql, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}

	var id int64
	if err := db.Conn.QueryRow(sql, args...).Scan(&id); err != nil {
		return nil, err
	}

	i := int(id)
	return &i, nil
}
