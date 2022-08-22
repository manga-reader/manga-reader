package database

import "fmt"

func (d *Database) Exec(cmd string) error {
	_, err := d.Instance.Exec(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute cmd '%s': %w", cmd, err)
	}
	return nil
}

func (d *Database) Insert(cmd string) error {
	_, err := d.Instance.Exec(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute cmd '%s': %w", cmd, err)
	}
	return nil
}

func (d *Database) InsertReturnID(cmd string) (DatabaseID, error) {
	lastInsertID := 0
	err := d.Instance.QueryRow(cmd).Scan(&lastInsertID)
	if err != nil {
		return 0, fmt.Errorf("failed to execute cmd '%s': %w", cmd, err)
	}
	return DatabaseID(lastInsertID), nil
}
