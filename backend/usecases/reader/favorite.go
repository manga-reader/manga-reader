package reader

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/utils"
)

func (r *Reader) GetFavorites(from, to int) ([]*database.ComicInfo, error) {
	q := fmt.Sprintf("SELECT comics.id, comics.name, comics.latest_volume, comics.updated_at "+
		"FROM favorite "+
		"INNER JOIN comics ON comics.id=favorite.comic_id "+
		"WHERE favorite.reader_id='%s';",
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

func (r *Reader) AddFavorite(comicID string) error {
	var err error
	q := fmt.Sprintf("SELECT id FROM comics WHERE comics.id='%s';", comicID)
	err = r.db.IsExist(q)
	if err != nil && err != utils.ErrNotFound {
		return fmt.Errorf("failed to find whether a note exist: %w", err)
	}
	if err == utils.ErrNotFound {
		// TODO add comics
	}
	cmd := fmt.Sprintf(`INSERT INTO 
	favorite ( 
		reader_id,
		comic_id
	)
    SELECT '%s', '%s' 
	WHERE NOT EXISTS (
    	SELECT 1 FROM favorite WHERE favorite.reader_id='%s' AND favorite.comic_id='%s'
	);`,
		r.ID, comicID,
		r.ID, comicID)
	return r.db.Exec(cmd)
}

func (r *Reader) DelFavorite(comicID string) error {
	cmd := fmt.Sprintf(`DELETE FROM favorite
	WHERE favorite.reader_id='%s' AND favorite.comic_id='%s';`,
		r.ID, comicID)
	return r.db.Exec(cmd)
}
