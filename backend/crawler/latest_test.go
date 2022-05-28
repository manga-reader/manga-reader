package crawler_test

import (
	"testing"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/stretchr/testify/assert"
)

func Test_GetPageLatestVolume(t *testing.T) {
	latestVol, err := crawler.GetPageLatestVolume("https://www.comicabc.com/html/131.html")
	assert.NoError(t, err)
	assert.Equal(t, "108話", latestVol)
}
