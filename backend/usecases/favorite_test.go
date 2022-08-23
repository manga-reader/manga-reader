package usecases

import (
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/stretchr/testify/require"
)

func Test_FavoriteAddGetDel(t *testing.T) {
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
	err = u.AddFavorite(reader.ID, testComicID)
	require.NoError(t, err)
	infos, err := u.GetFavorites(reader.ID, 0, 0)
	require.NoError(t, err)
	require.Len(t, infos, 1)
	require.Equal(t, testComicID, infos[0].ID)
	require.Equal(t, testName, infos[0].Name)
	require.Equal(t, testLatestVol, infos[0].LatestVolume)
	err = u.DelFavorite(reader.ID, testComicID)
	require.NoError(t, err)
	infos, err = u.GetFavorites(reader.ID, 0, 0)
	require.NoError(t, err)
	require.Len(t, infos, 0)
}
