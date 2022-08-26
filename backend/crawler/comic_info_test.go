package crawler_test

import (
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/stretchr/testify/require"
)

func Test_GetComicInfo(t *testing.T) {
	comicID := "131"
	name, latestVol, updatedAt, err := crawler.GetComicInfo(comicID)
	require.NoError(t, err)
	require.Equal(t, "鋼之鏈金術師", name)
	require.Equal(t, "108話", latestVol)
	require.Equal(t, time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC), *updatedAt)
}
