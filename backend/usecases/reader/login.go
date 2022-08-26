package reader

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/database"
)

func Login(db *database.Database, name string) (*Reader, error) {
	cmd := fmt.Sprintf(`INSERT INTO 
	readers (
		id, 
		created_at
	)
    SELECT '%s', NOW()
	WHERE NOT EXISTS (
    	SELECT 1 FROM readers WHERE id='%s'
	);`, name, name)
	err := db.Insert(cmd)
	if err != nil {
		return nil, err
	}
	return &Reader{
		ID: name,
		db: db,
	}, nil
}
