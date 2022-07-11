package reader

import (
	"encoding/json"
	"fmt"

	"github.com/manga-reader/manga-reader/backend/database"
)

func (r *Reader) GetFavoriteList(db *database.Database) ([]database.ComicInfo, error) {
	resRaws, err := db.ListRangeAll(database.GetUserFavoriteKey(r.ID))
	if err != nil {
		return nil, fmt.Errorf("can't get user %s's favorite list: %w", r.ID, err)
	}

	res := make([]database.ComicInfo, len(resRaws))
	for i, raw := range resRaws {
		var tmp database.ComicInfo
		err = json.Unmarshal([]byte(raw), &tmp)
		if err != nil {
			return nil, fmt.Errorf("can't unmarshal raw message  '%s' : %w", raw, err)
		}
		res[i] = tmp
	}

	return res, nil
}
