package user

import "github.com/manga-reader/manga-reader/backend/database"

type User struct {
	ID       string
	Database *database.Database
}
