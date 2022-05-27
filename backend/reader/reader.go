package reader

import "github.com/manga-reader/manga-reader/backend/database"

type Reader struct {
	ID       string
	Database *database.Database
}

func GetReader(id string, db *database.Database) *Reader {
	return &Reader{
		ID:       id,
		Database: db,
	}
}
