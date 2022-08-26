package scheduler

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/manga-reader/manga-reader/backend/database"
)

func UpdateFavoriteOrder(db *database.Database, userID string) error {
	ids, err := getFavoriteComicIDs(db, userID)
	if err != nil {
		return fmt.Errorf("err in calling getFavoriteComicIDs: %w", err)
	}

	favoriteComicInfos := make([]*database.ComicInfo, len(ids))
	for i, id := range ids {
		latestVol, updatedAt, err := crawler.GetPageLatestVolumeAndDate(ids[i])
		if err != nil {
			return fmt.Errorf("failed to get latest volume of comic: %v: %w", id, err)
		}

		comicInfo, err := getComicInfoByComicID(db, id)
		if err != nil {
			return fmt.Errorf("failed to call getComicInfoByComicID: %w", err)
		}

		if comicInfo.LatestVolume != latestVol {
			comicInfo.LatestVolume = latestVol
			comicInfo.UpdatedAt = *updatedAt
		}

		favoriteComicInfos[i] = comicInfo
	}

	err = db.Del(database.GetUserFavoriteKey(userID))
	if err != nil {
		return fmt.Errorf("failed to deleted outdated user favorite list of user: %s: %w", userID, err)
	}

	// sort UpdatedAt in ascending order (latest one in the end of slice)
	sort.Slice(favoriteComicInfos, func(i, j int) bool {
		return favoriteComicInfos[i].UpdatedAt.Before(favoriteComicInfos[j].UpdatedAt)
	})

	updatedComicIDs := make([]string, len(favoriteComicInfos))
	for i, info := range favoriteComicInfos {
		updatedComicIDs[i] = info.ID
	}

	err = db.ListPush(database.GetUserFavoriteKey(userID), updatedComicIDs)
	if err != nil {
		return fmt.Errorf("failed to push updated comic IDs to user favorite list of user: %s: %w", userID, err)
	}

	for _, info := range favoriteComicInfos {
		b, err := json.Marshal(info)
		if err != nil {
			return fmt.Errorf("failed to marshal comic info: %v to update: %w", info, err)
		}
		err = db.Set(info.ID, b)
		if err != nil {
			return fmt.Errorf("failed to update new comic info: %v: %w", info, err)
		}
	}

	return nil
}

func getFavoriteComicIDs(db *database.Database, userID string) ([]string, error) {
	ids, err := db.ListRangeAll(database.GetUserFavoriteKey(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to get favorite comic IDs of user %s: %w", userID, err)
	}
	return ids, nil
}

func getComicInfoByComicID(db *database.Database, id string) (*database.ComicInfo, error) {
	rawComicInfo, err := db.Get(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get comic info of id %s: %w", id, err)
	}

	var info database.ComicInfo
	err = json.Unmarshal([]byte(rawComicInfo), &info)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw comic info: %s: %w", info, err)
	}

	return &info, nil
}
