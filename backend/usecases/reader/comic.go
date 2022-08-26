package reader

import (
	"fmt"
	"time"
)

func (r *Reader) AddComic(comicID, name, latestVolume string, updatedAt time.Time) error {
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
	return r.db.Insert(cmd)
}
