package usecases

import (
	"fmt"
)

func (r *Reader) RecordSave(website WebsiteType, comicID, volume string, page int) error {
	cmd := fmt.Sprintf(`INSERT INTO 
	history (
		reader_id, 
		comic_id,
		volume,
		page,
		read_at
	)
	VALUES('%s', '%s', '%s', %d, NOW())
	ON CONFLICT (reader_id, comic_id)
	DO 
	   UPDATE SET volume='%s', page=%d;`,
		r.ID, comicID, volume, page,
		volume, page,
	)
	return r.db.Exec(cmd)
}

func (r *Reader) RecordLoad(website WebsiteType, comicID string) (string, int, error) {
	q := fmt.Sprintf("SELECT history.volume, history.page "+
		"FROM history "+
		"WHERE history.reader_id='%s';",
		r.ID)
	rows, err := r.db.Query(q)
	if err != nil {
		return "", 0, fmt.Errorf("failed to query: '%s': %w", q, err)
	}
	defer rows.Close()

	var vol string
	var page int
	for rows.Next() {
		err := rows.Scan(&vol, &page)
		if err != nil {
			return "", 0, fmt.Errorf("failed to scan response of query '%s': %w", q, err)
		}
	}

	if err = rows.Err(); err != nil {
		return "", 0, fmt.Errorf("failed to get response of query '%s': %w", q, err)
	}

	return vol, page, nil
}
