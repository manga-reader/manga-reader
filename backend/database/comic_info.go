package database

import "time"

type ComicInfo struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	LatestVolume string    `json:"latest_volume,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
