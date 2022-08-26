package usecases

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/stretchr/testify/require"
)

func TestDatabase_DropTables(t *testing.T) {
	type fields struct {
		host     string
		port     int
		user     string
		password string
		dbname   string
	}
	type wants struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		wants  wants
	}{
		{
			name: "successful",
			fields: fields{
				host:     database.Default_Host,
				port:     database.Default_Port,
				user:     database.Default_User,
				password: database.Default_Password,
				dbname:   database.Default_Dbname,
			},
			wants: wants{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := database.NewDatabase(
				tt.fields.host,
				tt.fields.port,
				tt.fields.user,
				tt.fields.password,
				tt.fields.dbname,
			)
			err := d.Connect()
			require.NoError(t, err)
			err = DropTables(d)
			require.NoError(t, err)

		})
	}
}
