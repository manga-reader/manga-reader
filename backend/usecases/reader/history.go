package reader

// func (r *Reader) GetHistoryList(db *database.Database) ([]database.ComicInfo, error) {
// 	historyIDs, err := db.ListRangeAll(database.GetUserHistoryKey(r.ID))
// 	if err != nil {
// 		return nil, fmt.Errorf("can't get user %s's history list: %w", r.ID, err)
// 	}

// 	res := make([]database.ComicInfo, len(historyIDs))
// 	for i, historyID := range historyIDs {
// 		infoRaw, err := db.Get(historyID)
// 		if err != nil {
// 			return nil, fmt.Errorf("can't get history comic: %s info: %w", historyID, err)
// 		}

// 		var tmp database.ComicInfo
// 		err = json.Unmarshal([]byte(infoRaw), &tmp)
// 		if err != nil {
// 			return nil, fmt.Errorf("can't unmarshal raw message '%s' : %w", infoRaw, err)
// 		}
// 		res[i] = tmp
// 	}

// 	return res, nil
// }

// func (r *Reader) AddNewHistory(db *database.Database, comicID string) error {
// 	comicName, err := crawler.GetComicName(comicID)
// 	if err != nil {
// 		return fmt.Errorf("can't get comic %s's name: %w", comicID, err)
// 	}

// 	latestVol, updatedAt, err := crawler.GetPageLatestVolumeAndDate(comicID)
// 	if err != nil {
// 		return fmt.Errorf("failed to get latest volume of comic: %v: %w", comicID, err)
// 	}

// 	err = db.ListPush(database.GetUserHistoryKey(r.ID), []string{comicID})
// 	if err != nil {
// 		return fmt.Errorf("failed to add comic id: %v to user history list: %w", comicID, err)
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

// func (r *Reader) DelHistory(db *database.Database, comicID string) error {
// 	err := db.ListRemoveElement(database.GetUserHistoryKey(r.ID), comicID)
// 	if err != nil {
// 		return fmt.Errorf("failed to remove comic id: %s from user: %s: %w", comicID, r.ID, err)
// 	}

// 	return nil
// }
