package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func (d *Database) Connect() error {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return fmt.Errorf("can't connect to database with '%s': %w", conn, err)
	}
	d.Instance = db
	return nil
}

func (d *Database) Ping() error {
	return d.Instance.Ping()
}
