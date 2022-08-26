package usecases

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/database"
)

// calling order matters, since it would affect the table relations
func DropTables(db *database.Database) error {
	var err error
	err = db.Exec("DROP TABLE history;")
	if err != nil {
		return fmt.Errorf("failed to drop table 'history': %w", err)
	}
	err = db.Exec("DROP TABLE favorite;")
	if err != nil {
		return fmt.Errorf("failed to drop table 'favorite': %w", err)
	}
	err = db.Exec("DROP TABLE comics;")
	if err != nil {
		return fmt.Errorf("failed to drop table 'comics': %w", err)
	}
	err = db.Exec("DROP TABLE readers;")
	if err != nil {
		return fmt.Errorf("failed to drop table 'readers': %w", err)
	}
	return nil
}
