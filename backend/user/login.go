package user

import "github.com/manga-reader/manga-reader/backend/database"

func Login(name string, database *database.Database) *User {
	return &User{
		ID:       name,
		Database: database,
	}
}
