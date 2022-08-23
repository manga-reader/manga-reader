package crawler_test

import (
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/stretchr/testify/assert"
)

func Test_GetComicInfo(t *testing.T) {
	comicID := "131"
	name, latestVol, updatedAt, err := crawler.GetComicInfo(comicID)
	assert.NoError(t, err)
	assert.Equal(t, "鋼之鏈金術師", name)
	assert.Equal(t, "108話", latestVol)
	assert.Equal(t, time.Date(2018, time.July, 23, 0, 0, 0, 0, time.Local), *updatedAt)
}
