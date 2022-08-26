package reader

import "github.com/manga-reader/manga-reader/backend/database"

type Reader struct {
	ID string
	db *database.Database
}

func GetReader(id string, db *database.Database) *Reader {
	return &Reader{
		ID: id,
		db: db,
	}
}
