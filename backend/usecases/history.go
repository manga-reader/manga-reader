package usecases

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/manga-reader/manga-reader/backend/utils"
)

func (u *Usecase) GetHistory(readerID string, from, to int) ([]*ComicInfo, error) {
	q := fmt.Sprintf("SELECT comics.id, comics.name, comics.latest_volume, comics.updated_at "+
		"FROM history "+
		"INNER JOIN comics ON comics.id=history.comic_id "+
		"WHERE history.reader_id='%s' "+
		"ORDER BY history.read_at DESC;",
		readerID)

	rows, err := u.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("failed to query cmd: '%s': %w", q, err)
	}
	defer rows.Close()

	var comicInfos []*ComicInfo
	for rows.Next() {
		var comicInfo ComicInfo
		err := rows.Scan(&comicInfo.ID, &comicInfo.Name, &comicInfo.LatestVolume, &comicInfo.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan response of query '%s': %w", q, err)
		}
		comicInfos = append(comicInfos, &comicInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get response of query '%s': %w", q, err)
	}

	return comicInfos, nil
}

func (u *Usecase) AddHistory(readerID, comicID string) error {
	var err error
	q := fmt.Sprintf("SELECT id FROM comics WHERE comics.id='%s';", comicID)
	err = u.db.IsExist(q)
	if err != nil && err != utils.ErrNotFound {
		return fmt.Errorf("failed to find whether a note exist: %w", err)
	}
	if err == utils.ErrNotFound {
		name, latestVol, updatedAt, err := crawler.GetComicInfo(comicID)
		if err != nil && err != utils.ErrNotFound {
			return fmt.Errorf("failed to get comic info of %s by crawler: %w", comicID, err)
		}
		err = u.AddComic(comicID, name, latestVol, updatedAt)
		if err != nil && err != utils.ErrNotFound {
			return fmt.Errorf("failed to add comic with comic_id: %s, name: %s, latest volume: %s, updated at: %s: %w", comicID, name, latestVol, updatedAt, err)
		}
	}
	cmd := fmt.Sprintf(`INSERT INTO 
	history ( 
		reader_id,
		comic_id,
		read_at
	)
    SELECT '%s', '%s', NOW()
	WHERE NOT EXISTS (
    	SELECT 1 FROM history WHERE history.reader_id='%s' AND history.comic_id='%s'
	);`,
		readerID, comicID,
		readerID, comicID)
	return u.db.Exec(cmd)
}

func (u *Usecase) DelHistory(readerID, comicID string) error {
	cmd := fmt.Sprintf(`DELETE FROM history
	WHERE history.reader_id='%s' AND history.comic_id='%s';`,
		readerID, comicID)
	return u.db.Exec(cmd)
}
