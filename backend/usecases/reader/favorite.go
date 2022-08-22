package reader

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/database"
)

func (r *Reader) GetFavoriteList(readerID string, from, to int) ([]*database.ComicInfo, error) {
	q := fmt.Sprintf("SELECT comics.id, comics.name, comics.latest_volume, comics.updated_at "+
		"FROM favorite "+
		"INNER JOIN comics.id=favorite.comic_id "+
		"WHERE favorite.reader_id='%s';",
		readerID)
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("failed to query cmd: '%s': %w", q, err)
	}
	defer rows.Close()

	var comicInfos []*database.ComicInfo
	for rows.Next() {
		var comicInfo database.ComicInfo
		err := rows.Scan(&comicInfo.ID, &comicInfo.Name, &comicInfo.LatestVolume, &comicInfo.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan response of query '%s': %w", q, err)
		}
		comicInfos = append(comicInfos)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get response of query '%s': %w", q, err)
	}

	return comicInfos, nil
}

// func (r *Reader) AddNewFavorite(db *database.Database, comicID string) error {
// 	comicName, err := crawler.GetComicName(comicID)
// 	if err != nil {
// 		return fmt.Errorf("can't get comic %s's name: %w", comicID, err)
// 	}

// 	latestVol, updatedAt, err := crawler.GetPageLatestVolumeAndDate(comicID)
// 	if err != nil {
// 		return fmt.Errorf("failed to get latest volume of comic: %v: %w", comicID, err)
// 	}

// 	err = db.ListPush(database.GetUserFavoriteKey(r.ID), []string{comicID})
// 	if err != nil {
// 		return fmt.Errorf("failed to add comic id: %v to user favorite list: %w", comicID, err)
// 	}

// 	info := database.ComicInfo{
// 		ID:           comicID,
// 		Name:         comicName,
// 		LatestVolume: latestVol,
// 		UpdatedAt:    *updatedAt,
// 	}
// 	b, err := json.Marshal(&info)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal comic info: %v to update: %w", info, err)
// 	}

// 	return db.Set(comicID, b)
// }

// func (r *Reader) DelFavorite(db *database.Database, comicID string) error {
// 	err := db.ListRemoveElement(database.GetUserFavoriteKey(r.ID), comicID)
// 	if err != nil {
// 		return fmt.Errorf("failed to remove comic id: %s from user: %s: %w", comicID, r.ID, err)
// 	}

// 	return nil
// }
