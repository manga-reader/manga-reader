package usecases

import "github.com/manga-reader/manga-reader/backend/database"

type Reader struct {
	ID string
}

type WebsiteType int

const (
	Unspecified WebsiteType = iota
	Website_8comic
)

type Usecase struct {
	db *database.Database
}

func NewUsecase(db *database.Database) *Usecase {
	return &Usecase{
		db: db,
	}
}
