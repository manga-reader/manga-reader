package router_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manga-reader/manga-reader/backend/router"
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
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println("w.Body.String(): ", w.Body.String())
}
