package router_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manga-reader/manga-reader/backend/config"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/router/handler/process"
	"github.com/manga-reader/manga-reader/backend/router/handler/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func Test_ProcessSaveAndLoad(t *testing.T) {
	db := database.Connect(config.Cfg.Redis.ServerAddr, config.Cfg.Redis.Password, config.Cfg.Redis.DBIndex)
	router := router.SetupRouter(
		&router.Params{db},
	)

	wSave := httptest.NewRecorder()
	comicID := "7"
	vol := "10"
	page := 3

	reqSave, _ := http.NewRequest("GET", "/process/save", nil)
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
	reqLoad, _ := http.NewRequest("GET", "/process/load", nil)
	q = reqLoad.URL.Query()
	q.Add(handler.HeaderComicID, comicID)
	reqLoad.URL.RawQuery = q.Encode()
	reqLoad.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.N3sjQ9IX8ipYMA9bxT4PyvSTRYLIKFwvkYu-hnNVqvM")
	router.ServeHTTP(wLoad, reqLoad)

	var processLoadRes process.ProcessLoadRes
	err := json.Unmarshal(wLoad.Body.Bytes(), &processLoadRes)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, wLoad.Code)
	assert.Equal(t, vol, processLoadRes.Volume)
	assert.Equal(t, page, processLoadRes.Page)

}
