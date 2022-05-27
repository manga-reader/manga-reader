package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/sirupsen/logrus"
)

type UserLoginRes struct {
	Token string `json:"token,omitempty"`
}

func UserLogin(c *gin.Context) {
	userID := getUserLoginQueryParams(c)
	tokenString, err := auth.GenerateNewToken(userID)
	if err != nil {
		logrus.Errorf("failed to generate JWT string: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to generate JWT string",
		})
	}

	res := &UserLoginRes{
		Token: tokenString,
	}

	c.JSON(http.StatusOK, res)
}

func getUserLoginQueryParams(c *gin.Context) string {
	if c.Query(handler.HeaderUserID) == "" {
		logrus.Errorf("user id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "user id is not given"})
	}

	userID := c.Query(handler.HeaderUserID)
	return userID
}
