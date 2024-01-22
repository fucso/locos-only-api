package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

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
