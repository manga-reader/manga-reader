package router_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/router/handler/user"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.-aDZ-M5kwECjejs1jUiMcuyaKMYjAxknyObs4i4elcY", res.Token)
}

func Test_ProcessSave(t *testing.T) {
	router := router.SetupRouter(
		&router.Params{},
	)

	w := httptest.NewRecorder()
	comicID := "7"
	vol := "10"
	page := "3"

	req, _ := http.NewRequest("GET", "/process/save", nil)
	q := req.URL.Query()
	q.Add(handler.HeaderComicID, comicID)
	q.Add(handler.HeaderVolume, vol)
	q.Add(handler.HeaderPage, page)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImpvaG4ifQ.-aDZ-M5kwECjejs1jUiMcuyaKMYjAxknyObs4i4elcY")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, handler.ResponseOK, w.Body.String())
}
