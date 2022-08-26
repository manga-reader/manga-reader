package scheduler

import (
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/stretchr/testify/require"
)

func TestUpdateFavorites(t *testing.T) {
	testUserID := "john"
	comicInfosBefore := []*usecases.ComicInfo{
		{
			ID:           "3654",
			Name:         "妖精的尾巴",
			LatestVolume: "540話",
			UpdatedAt:    time.Date(2017, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           "131",
			Name:         "鋼之鏈金術師",
			LatestVolume: "20卷",
			UpdatedAt:    time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           "9337",
			Name:         "食戟之靈",
			LatestVolume: "315話",
			UpdatedAt:    time.Date(2019, time.December, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	comicInfosAfter := []*usecases.ComicInfo{
		{
			ID:           "9337",
			Name:         "食戟之靈",
			LatestVolume: "315話",
			UpdatedAt:    time.Date(2020, time.May, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           "3654",
			Name:         "妖精的尾巴",
			LatestVolume: "545話 無法取代的伙伴們",
			UpdatedAt:    time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           "131",
			Name:         "鋼之鏈金術師",
			LatestVolume: "108話",
			UpdatedAt:    time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
	}

	db := database.NewDatabase(
		database.Default_Host,
		database.Default_Port,
		database.Default_User,
		database.Default_Password,
		database.Default_Dbname,
	)
	err := db.Connect()
	require.NoError(t, err)
	u := usecases.NewUsecase(db)
	require.NotNil(t, u)
	s := NewScheduler(u)
	require.NotNil(t, s)
	reader, err := u.Login(testUserID)
	require.NoError(t, err)

	for _, info := range comicInfosBefore {
		err = u.AddFavorite(reader.ID, info.ID)
		require.NoError(t, err)
	}

	err = s.UpdateFavorites(testUserID)
	require.NoError(t, err)

	infos, err := u.GetFavorites(reader.ID, 0, 0)
	require.NoError(t, err)
	require.Equal(t, comicInfosAfter, infos)
}
