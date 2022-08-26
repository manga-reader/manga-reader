package router_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/config"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/router/handler/record"
	"github.com/manga-reader/manga-reader/backend/router/handler/user"
	"github.com/manga-reader/manga-reader/backend/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	db := database.Connect(config.Cfg.Redis.ServerAddr, config.Cfg.Redis.Password, config.Cfg.Redis.DBIndex)
	err := db.FlushAll()
	if err != nil {
		panic(err)
	}
}

func Test_HealthPing(t *testing.T) {
	router := router.SetupRouter(
		&router.Params{},
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func Test_UserLogin(t *testing.T) {
	router := router.SetupRouter(
		&router.Params{},
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/login", nil)
	q := req.URL.Query()
	q.Add(handler.HeaderUserID, "john")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var res user.UserLoginRes
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM", res.Token)
}

func Test_RecordSaveAndLoad(t *testing.T) {
	db := database.Connect(config.Cfg.Redis.ServerAddr, config.Cfg.Redis.Password, config.Cfg.Redis.DBIndex)
	router := router.SetupRouter(
		&router.Params{db},
	)

	wSave := httptest.NewRecorder()
	comicID := "7"
	vol := "10"
	page := 3

	reqSave, _ := http.NewRequest("GET", "/record/save", nil)
	q := reqSave.URL.Query()
	q.Add(handler.HeaderComicID, comicID)
	q.Add(handler.HeaderVolume, vol)
	q.Add(handler.HeaderPage, fmt.Sprint(page))
	reqSave.URL.RawQuery = q.Encode()
	reqSave.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(wSave, reqSave)

	assert.Equal(t, http.StatusOK, wSave.Code)
	assert.Equal(t, handler.ResponseOK, wSave.Body.String())

	wLoad := httptest.NewRecorder()
	reqLoad, _ := http.NewRequest("GET", "/record/load", nil)
	q = reqLoad.URL.Query()
	q.Add(handler.HeaderComicID, comicID)
	reqLoad.URL.RawQuery = q.Encode()
	reqLoad.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(wLoad, reqLoad)

	var recordLoadRes record.RecordLoadRes
	err := json.Unmarshal(wLoad.Body.Bytes(), &recordLoadRes)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, wLoad.Code)
	assert.Equal(t, vol, recordLoadRes.Volume)
	assert.Equal(t, page, recordLoadRes.Page)
}

func Test_FavoriteAddGetDel(t *testing.T) {
	// TODO
}

func Test_HistoryGet(t *testing.T) {
	testIDs := []string{"123", "456", "789"}
	testComicInfos := []*database.ComicInfo{
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
	db := database.Connect(config.Cfg.Redis.ServerAddr, config.Cfg.Redis.Password, config.Cfg.Redis.DBIndex)
	err := db.ListPush(database.GetUserHistoryKey("john"), utils.ReverseStringSlice(testIDs))
	require.NoError(t, err)

	for i := range testIDs {
		b, err := json.Marshal(testComicInfos[i])
		require.NoError(t, err)
		err = db.Set(testIDs[i], b)
		require.NoError(t, err)
	}

	router := router.SetupRouter(
		&router.Params{db},
	)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/history", nil)
	require.NoError(t, err)

	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(w, req)

	var res []*database.ComicInfo
	err = json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.Equal(t, testComicInfos, res)
}
