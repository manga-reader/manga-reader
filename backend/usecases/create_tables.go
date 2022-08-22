package usecases

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/database"
)

// calling order matters, since it would affect the table relations
func CreateTables(db *database.Database) error {
	var err error
	err = db.Exec(CREATE_READERS_TABLE)
	if err != nil {
		return fmt.Errorf("failed to create table 'readers': %w", err)
	}
	err = db.Exec(CREATE_COMICS_TABLE)
	if err != nil {
		return fmt.Errorf("failed to create table 'comics': %w", err)
	}
	err = db.Exec(CREATE_FAVORITE_TABLE)
	if err != nil {
		return fmt.Errorf("failed to create table 'favorite': %w", err)
	}
	err = db.Exec(CREATE_HISTORY_TABLE)
	if err != nil {
		return fmt.Errorf("failed to create table 'history': %w", err)
	}
	return nil
}
