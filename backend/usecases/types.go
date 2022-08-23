package usecases

import "github.com/manga-reader/manga-reader/backend/database"

type Reader struct {
	ID string
	db *database.Database
}

type WebsiteType int

const (
	Unspecified WebsiteType = iota
	Website_8comic
)
