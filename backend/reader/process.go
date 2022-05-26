package reader

import "fmt"

func (r *Reader) ProcessSave(website WebsiteType, comicID, volume, page string) error {
	key := fmt.Sprintf("%s:%d:%s", r.ID, website, comicID)
	record := fmt.Sprintf("%s:%s", volume, page)
	return r.Database.Set(key, record)
}

func (r *Reader) ProcessLoad(website WebsiteType, comicID string) (string, error) {
	key := fmt.Sprintf("%s:%d:%s", r.ID, website, comicID)
	return r.Database.Get(key)
}
