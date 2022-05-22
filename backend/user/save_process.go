package user

import "fmt"

func (u *User) SaveProcess(website WebsiteType, comicID, volume, page string) error {
	key := fmt.Sprintf("%s:%d:%s", u.ID, website, comicID)
	record := fmt.Sprintf("%s:%s", volume, page)
	return u.Database.Set(key, record)
}
