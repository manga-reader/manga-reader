package usecases

import (
	"testing"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/stretchr/testify/require"
)

func Test_RecordSaveLoad(t *testing.T) {
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
	testVol := "47"
	testPage := 25
	err = u.RecordSave(Website_8comic, reader.ID, testComicID, testVol, testPage)
	require.NoError(t, err)
	vol, page, err := u.RecordLoad(Website_8comic, reader.ID, testComicID)
	require.NoError(t, err)
	require.Equal(t, testVol, vol)
	require.Equal(t, testPage, page)
}
