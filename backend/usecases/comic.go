package usecases

import (
	"fmt"
	"time"
)

type ComicInfo struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	LatestVolume string    `json:"latest_volume,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (u *Usecase) AddComic(comicID, name, latestVolume string, updatedAt *time.Time) error {
	t := updatedAt.Format(time.RFC3339)
	cmd := fmt.Sprintf(`INSERT INTO 
	comics (
		id,
		name,
		latest_volume,
		updated_at
	)
    SELECT '%s', '%s', '%s', '%s'
	WHERE NOT EXISTS (
    	SELECT 1 FROM comics WHERE id='%s'
	);`,
		comicID, name, latestVolume, t, comicID)
	return u.db.Insert(cmd)
}

func (u *Usecase) GetComicByID(comicID string) (*ComicInfo, error) {
	q := fmt.Sprintf("SELECT id, name, latest_volume, updated_at "+
		"FROM comics "+
		"WHERE id='%s';",
		comicID)

	rows, err := u.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("failed to query cmd: '%s': %w", q, err)
	}
	defer rows.Close()

	var comicInfo ComicInfo
	for rows.Next() {
		err := rows.Scan(&comicInfo.ID, &comicInfo.Name, &comicInfo.LatestVolume, &comicInfo.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan response of query '%s': %w", q, err)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get response of query '%s': %w", q, err)
	}

	return &comicInfo, nil
}

func (u *Usecase) UpdateComic(comicID string, latestVolume string, updatedAt *time.Time) error {
	cmd := fmt.Sprintf("UPDATE comics SET latest_volume='%s', updated_at='%s';", latestVolume, updatedAt)
	return u.db.Exec(cmd)
}

func (u *Usecase) DelComic(comicID string) error {
	cmd := fmt.Sprintf(`DELETE FROM comics WHERE id='%s';`, comicID)
	return u.db.Exec(cmd)
}
