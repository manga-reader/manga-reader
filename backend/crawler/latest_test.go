package crawler_test

import (
	"testing"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/stretchr/testify/assert"
)

func Test_GetPageLatestVolume(t *testing.T) {
	comicID := "131"
	latestVol, err := crawler.GetPageLatestVolume(comicID)
	assert.NoError(t, err)
	assert.Equal(t, "108話", latestVol)
}
