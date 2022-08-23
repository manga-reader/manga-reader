package usecases

import (
	"fmt"
)

func (u *Usecase) Login(name string) (*Reader, error) {
	cmd := fmt.Sprintf(`INSERT INTO 
	readers (
		id, 
		created_at
	)
    SELECT '%s', NOW()
	WHERE NOT EXISTS (
    	SELECT 1 FROM readers WHERE id='%s'
	);`, name, name)
	err := u.db.Insert(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to add new reader: %s: %w", name, err)
	}
	return &Reader{
		ID: name,
	}, nil
}
