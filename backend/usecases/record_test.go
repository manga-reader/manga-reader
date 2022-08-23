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
	reader, err := Login(d, "john")
	require.NoError(t, err)
	require.NotNil(t, reader)
	testComicID := "123"
	testVol := "47"
	testPage := 25
	err = reader.RecordSave(Website_8comic, testComicID, testVol, testPage)
	require.NoError(t, err)
	vol, page, err := reader.RecordLoad(Website_8comic, testComicID)
	require.NoError(t, err)
	require.Equal(t, testVol, vol)
	require.Equal(t, testPage, page)
}
