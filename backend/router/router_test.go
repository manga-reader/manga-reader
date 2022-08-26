package router_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/router/handler/record"
	"github.com/manga-reader/manga-reader/backend/router/handler/user"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableQuote: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
}

func Test_HealthPing(t *testing.T) {
	router := router.SetupRouter(
		&router.Params{},
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/ping", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "pong", w.Body.String())
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

	require.Equal(t, http.StatusOK, w.Code)
	var res user.UserLoginRes
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM", res.Token)
}

func Test_RecordSaveAndLoad(t *testing.T) {
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
	router := router.SetupRouter(
		&router.Params{u},
	)

	wNew := httptest.NewRecorder()
	comicID := "131"
	vol := "90話"

	reqNew, _ := http.NewRequest("GET", "/record/new", nil)
	q := reqNew.URL.Query()
	q.Add(handler.HeaderComicID, comicID)
	q.Add(handler.HeaderVolume, vol)
	reqNew.URL.RawQuery = q.Encode()
	reqNew.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(wNew, reqNew)

	require.Equal(t, http.StatusOK, wNew.Code)
	require.Equal(t, handler.ResponseOK, wNew.Body.String())

	wSave := httptest.NewRecorder()
	vol = "96話"
	page := 3

	reqSave, _ := http.NewRequest("GET", "/record/save", nil)
	q = reqSave.URL.Query()
	q.Add(handler.HeaderComicID, comicID)
	q.Add(handler.HeaderVolume, vol)
	q.Add(handler.HeaderPage, fmt.Sprint(page))
	reqSave.URL.RawQuery = q.Encode()
	reqSave.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(wSave, reqSave)

	require.Equal(t, http.StatusOK, wSave.Code)
	require.Equal(t, handler.ResponseOK, wSave.Body.String())

	wLoad := httptest.NewRecorder()
	reqLoad, _ := http.NewRequest("GET", "/record/load", nil)
	q = reqLoad.URL.Query()
	q.Add(handler.HeaderComicID, comicID)
	reqLoad.URL.RawQuery = q.Encode()
	reqLoad.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(wLoad, reqLoad)

	var recordLoadRes record.RecordLoadRes
	err = json.Unmarshal(wLoad.Body.Bytes(), &recordLoadRes)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, wLoad.Code)
	require.Equal(t, vol, recordLoadRes.Volume)
	require.Equal(t, page, recordLoadRes.Page)
}

func Test_FavoriteAddGetDel(t *testing.T) {
	var err error
	testComicInfos := []*usecases.ComicInfo{
		{
			ID:           "3654",
			Name:         "妖精的尾巴",
			LatestVolume: "545話 無法取代的伙伴們",
			UpdatedAt:    time.Date(2020, time.May, 01, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           "131",
			Name:         "鋼之鏈金術師",
			LatestVolume: "108話",
			UpdatedAt:    time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           "9337",
			Name:         "食戟之靈",
			LatestVolume: "315話",
			UpdatedAt:    time.Date(2020, time.May, 12, 0, 0, 0, 0, time.UTC),
		},
	}
	testComicInfosRes := []*usecases.ComicInfo{
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
			UpdatedAt:    time.Date(2020, time.May, 01, 0, 0, 0, 0, time.UTC),
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
	err = db.Connect()
	require.NoError(t, err)
	u := usecases.NewUsecase(db)
	require.NotNil(t, u)
	_, err = u.Login("john")
	require.NoError(t, err)
	router := router.SetupRouter(
		&router.Params{u},
	)

	for _, info := range testComicInfos {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/user/favorite", nil)
		require.NoError(t, err)
		q := req.URL.Query()
		q.Add(handler.HeaderComicID, info.ID)
		req.URL.RawQuery = q.Encode()
		req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
		router.ServeHTTP(w, req)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/favorite", nil)
	require.NoError(t, err)
	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(w, req)
	var res []*usecases.ComicInfo
	err = json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.Equal(t, testComicInfosRes, res)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", "/user/favorite", nil)
	require.NoError(t, err)
	q := req.URL.Query()
	q.Add(handler.HeaderComicID, testComicInfosRes[0].ID)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/user/favorite", nil)
	require.NoError(t, err)
	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(w, req)
	err = json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.Equal(t, testComicInfosRes[1:], res)
}

func Test_HistoryGet(t *testing.T) {
	testComicInfos := []*usecases.ComicInfo{
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
		{
			ID:           "9337",
			Name:         "食戟之靈",
			LatestVolume: "315話",
			UpdatedAt:    time.Date(2020, time.May, 12, 0, 0, 0, 0, time.UTC),
		},
	}
	testComicInfosRes := []*usecases.ComicInfo{
		{
			ID:           "131",
			Name:         "鋼之鏈金術師",
			LatestVolume: "108話",
			UpdatedAt:    time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC),
		},
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
	reader, err := u.Login("john")
	require.NoError(t, err)

	for _, comicInfo := range testComicInfos {
		err = u.AddHistory(reader.ID, comicInfo.ID)
		require.NoError(t, err)
	}

	router := router.SetupRouter(
		&router.Params{u},
	)

	for i := range testComicInfosRes {
		err = u.RecordSave(usecases.Website_8comic, reader.ID, testComicInfosRes[len(testComicInfosRes)-i-1].ID, "test_empty", 0)
		require.NoError(t, err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/history", nil)
	require.NoError(t, err)

	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(w, req)

	var res []*usecases.ComicInfo
	err = json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.Equal(t, testComicInfosRes, res)
}
