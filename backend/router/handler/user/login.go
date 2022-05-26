package reader

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
)

type UserLoginRes struct {
	Token string `json:"token,omitempty"`
}

func UserLogin(c *gin.Context) {
	tokenString, err := auth.GenerateNewToken(c.Param("id"))
	if err != nil {
		panic(err)
	}

	res := &UserLoginRes{
		Token: tokenString,
	}

	c.JSON(http.StatusOK, res)
}
