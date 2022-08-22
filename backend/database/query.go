package database

import (
	"database/sql"
	"fmt"

	"github.com/manga-reader/manga-reader/backend/utils"
)

func (d *Database) Query(q string) (*sql.Rows, error) {
	rows, err := d.Instance.Query(q)
	if err != nil {
		return nil, fmt.Errorf("failed to query by '%s': %w", q, err)
	}
	return rows, nil
}

func (d *Database) IsExist(q string) error {
	id := 0
	err := d.Instance.QueryRow(q).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("failed to check existence by q: '%s': %w", q, err)
		}
		return utils.ErrNotFound
	}
	return nil
}
