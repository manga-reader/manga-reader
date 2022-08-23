package scheduler

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateFavoriteOrder(t *testing.T) {
	testUserID := "test_user_id"
	keyFavorite := fmt.Sprintf("%s:favorite", testUserID)
	db := database.NewDatabase(
		database.Default_Host,
		database.Default_Port,
		database.Default_User,
		database.Default_Password,
		database.Default_Dbname,
	)
	err := db.Connect()
	require.NoError(t, err)

	err = db.Del(keyFavorite)
	require.NoError(t, err)

	comicInfosBefore := []*usecases.ComicInfo{
		{
			// 妖精的尾巴
			ID:           "3654",
			LatestVolume: "540話",
			UpdatedAt:    time.Date(2017, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			// 鋼之鏈金術師
			ID:           "131",
			LatestVolume: "20卷",
			UpdatedAt:    time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			// 曾為我兄者
			ID:           "19503",
			LatestVolume: "01",
			UpdatedAt:    time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
	}

	comicInfosAfter := []*usecases.ComicInfo{
		{
			// 曾為我兄者
			ID:           "19503",
			LatestVolume: "09.5",
			UpdatedAt:    time.Date(2022, time.August, 18, 0, 0, 0, 0, time.UTC),
		},
		{
			// 妖精的尾巴
			ID:           "3654",
			LatestVolume: "545話 無法取代的伙伴們",
			UpdatedAt:    time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			// 鋼之鏈金術師
			ID:           "131",
			LatestVolume: "108話",
			UpdatedAt:    time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, info := range comicInfosBefore {
		infoByte, err := json.Marshal(&info)
		require.NoError(t, err)
		err = db.ListPush(keyFavorite, []string{info.ID})
		require.NoError(t, err)
		err = db.Set(info.ID, infoByte)
		require.NoError(t, err)
	}

	err = UpdateFavoriteOrder(db, testUserID)
	require.NoError(t, err)

	updatedIDs, err := db.ListRangeAll(keyFavorite)
	require.NoError(t, err)
	for i, id := range updatedIDs {
		infoRaw, err := db.Get(id)
		require.NoError(t, err)
		var info usecases.ComicInfo
		err = json.Unmarshal([]byte(infoRaw), &info)
		require.NoError(t, err)
		assert.Equal(t, comicInfosAfter[i], &info)
	}

	err = db.Del(keyFavorite)
	require.NoError(t, err)
}
