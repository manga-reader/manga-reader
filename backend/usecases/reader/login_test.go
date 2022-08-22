package reader

import (
	"testing"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
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
}
