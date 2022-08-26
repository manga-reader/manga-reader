package reader

import (
	"encoding/json"
	"fmt"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/manga-reader/manga-reader/backend/database"
)

func (r *Reader) GetFavoriteList(db *database.Database) ([]database.ComicInfo, error) {
	favoriteIDs, err := db.ListRangeAll(database.GetUserFavoriteKey(r.ID))
	if err != nil {
		return nil, fmt.Errorf("can't get user %s's favorite list: %w", r.ID, err)
	}

	res := make([]database.ComicInfo, len(favoriteIDs))
	for i, favoriteID := range favoriteIDs {
		infoRaw, err := db.Get(favoriteID)
		if err != nil {
			return nil, fmt.Errorf("can't get favorite comic: %s info: %w", favoriteID, err)
		}

		var tmp database.ComicInfo
		err = json.Unmarshal([]byte(infoRaw), &tmp)
		if err != nil {
			return nil, fmt.Errorf("can't unmarshal raw message '%s' : %w", infoRaw, err)
		}
		res[i] = tmp
	}

	return res, nil
}

func (r *Reader) AddNewFavorite(db *database.Database, comicID string) error {
	comicName, err := crawler.GetComicName(comicID)
	if err != nil {
		return fmt.Errorf("can't get comic %s's name: %w", comicID, err)
	}

	latestVol, updatedAt, err := crawler.GetPageLatestVolumeAndDate(comicID)
	if err != nil {
		return fmt.Errorf("failed to get latest volume of comic: %v: %w", comicID, err)
	}

	err = db.ListPush(database.GetUserFavoriteKey(r.ID), []string{comicID})
	if err != nil {
		return fmt.Errorf("failed to add comic id: %v to user favorite list: %w", comicID, err)
	}

	info := database.ComicInfo{
		ID:           comicID,
		Name:         comicName,
		LatestVolume: latestVol,
		UpdatedAt:    *updatedAt,
	}
	b, err := json.Marshal(&info)
	if err != nil {
		return fmt.Errorf("failed to marshal comic info: %v to update: %w", info, err)
	}

	return db.Set(comicID, b)
}

func (r *Reader) DelFavorite(db *database.Database, comicID string) error {
	err := db.ListRemoveElement(database.GetUserFavoriteKey(r.ID), comicID)
	if err != nil {
		return fmt.Errorf("failed to remove comic id: %s from user: %s: %w", comicID, r.ID, err)
	}

	return nil
}
