package reader

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/utils"
)

func (r *Reader) GetHistory(from, to int) ([]*database.ComicInfo, error) {
	q := fmt.Sprintf("SELECT comics.id, comics.name, comics.latest_volume, comics.updated_at "+
		"FROM history "+
		"INNER JOIN comics ON comics.id=history.comic_id "+
		"WHERE history.reader_id='%s';",
		r.ID)

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
		comicInfos = append(comicInfos, &comicInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get response of query '%s': %w", q, err)
	}

	return comicInfos, nil
}

func (r *Reader) AddHistory(comicID string) error {
	var err error
	q := fmt.Sprintf("SELECT id FROM comics WHERE comics.id='%s';", comicID)
	err = r.db.IsExist(q)
	if err != nil && err != utils.ErrNotFound {
		return fmt.Errorf("failed to find whether a note exist: %w", err)
	}
	if err == utils.ErrNotFound {
		name, latestVol, updatedAt, err := crawler.GetComicInfo(comicID)
		if err != nil && err != utils.ErrNotFound {
			return fmt.Errorf("failed to get comic info of %s by crawler: %w", comicID, err)
		}
		err = r.AddComic(comicID, name, latestVol, *updatedAt)
		if err != nil && err != utils.ErrNotFound {
			return fmt.Errorf("failed to add comic with comic_id: %s, name: %s, latest volume: %s, updated at: %s: %w", comicID, name, latestVol, updatedAt, err)
		}
	}
	cmd := fmt.Sprintf(`INSERT INTO 
	history ( 
		reader_id,
		comic_id
	)
    SELECT '%s', '%s' 
	WHERE NOT EXISTS (
    	SELECT 1 FROM history WHERE history.reader_id='%s' AND history.comic_id='%s'
	);`,
		r.ID, comicID,
		r.ID, comicID)
	return r.db.Exec(cmd)
}

func (r *Reader) DelHistory(comicID string) error {
	cmd := fmt.Sprintf(`DELETE FROM history
	WHERE history.reader_id='%s' AND history.comic_id='%s';`,
		r.ID, comicID)
	return r.db.Exec(cmd)
}
