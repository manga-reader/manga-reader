package reader

import (
	"fmt"
	"strconv"
	"strings"
)

func (r *Reader) ProcessSave(website WebsiteType, comicID, volume, page string) error {
	key := fmt.Sprintf("%s:%d:%s", r.ID, website, comicID)
	record := fmt.Sprintf("%s:%s", volume, page)
	return r.Database.Set(key, record)
}

func (r *Reader) ProcessLoad(website WebsiteType, comicID string) (string, int, error) {
	key := fmt.Sprintf("%s:%d:%s", r.ID, website, comicID)
	val, err := r.Database.Get(key)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get process by key %v: %w", key, err)
	}
	vals := strings.Split(val, ":")
	vol := vals[0]
	page, err := strconv.Atoi(vals[1])
	if err != nil {
		return "", 0, fmt.Errorf("failed to get process by val %v: %w", val, err)
	}
	return vol, page, nil
}
