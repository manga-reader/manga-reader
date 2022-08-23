package scheduler

import (
	"fmt"

	"github.com/manga-reader/manga-reader/backend/crawler"
)

func (s *Scheduler) UpdateFavorites(readerID string) error {
	comicInfos, err := s.usecase.GetFavorites(readerID, 0, 0)
	if err != nil {
		return fmt.Errorf("err in calling getFavoriteComicIDs: %w", err)
	}

	for _, comicInfo := range comicInfos {
		_, latestVol, updatedAt, err := crawler.GetComicInfo(comicInfo.ID)
		if err != nil {
			return fmt.Errorf("failed to get latest volume of comic: %s: %w", comicInfo.ID, err)
		}

		if comicInfo.LatestVolume != latestVol {
			comicInfo.LatestVolume = latestVol
			comicInfo.UpdatedAt = *updatedAt
			err = s.usecase.UpdateComic(comicInfo.ID, latestVol, updatedAt)
			if err != nil {
				return fmt.Errorf("failed to update comic: %s with latestVol: %s, updatedAt: %s: %w", comicInfo.ID, latestVol, updatedAt, err)
			}
		}
	}
	return nil
}
