package crawler_test

import (
	"testing"

	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/stretchr/testify/assert"
)

func Test_GetGetComicName(t *testing.T) {
	comicID := "131"
	name, err := crawler.GetComicName(comicID)
	assert.NoError(t, err)
	assert.Equal(t, "鋼之鏈金術師", name)
}
