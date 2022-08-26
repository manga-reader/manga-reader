package usecases

import (
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/stretchr/testify/require"
)

func Test_AddComic(t *testing.T) {
	d := database.NewDatabase(
		database.Default_Host,
		database.Default_Port,
		database.Default_User,
		database.Default_Password,
		database.Default_Dbname,
	)
	err := d.Connect()
	require.NoError(t, err)
	u := NewUsecase(d)
	require.NotNil(t, u)
	reader, err := u.Login("john")
	require.NoError(t, err)
	require.NotNil(t, reader)
	testComicID := "123"
	testName := "test_comic_name"
	testLatestVol := "43"
	testUpdatedAt := time.Now()
	err = u.AddComic(testComicID, testName, testLatestVol, &testUpdatedAt)
	require.NoError(t, err)
	comicInfo, err := u.GetComicByID(testComicID)
	require.NoError(t, err)
	require.Equal(t, testComicID, comicInfo.ID)
	require.Equal(t, testName, comicInfo.Name)
	require.Equal(t, testLatestVol, comicInfo.LatestVolume)
	require.Equal(t, testUpdatedAt, comicInfo.UpdatedAt)
	err = u.DelComic(testComicID)
	require.NoError(t, err)
	_, err = u.GetComicByID(testComicID)
	require.NoError(t, err)
}
