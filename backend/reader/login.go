package reader

import "github.com/manga-reader/manga-reader/backend/database"

func Login(name string, database *database.Database) *Reader {
	return &Reader{
		ID:       name,
		Database: database,
	}
}
