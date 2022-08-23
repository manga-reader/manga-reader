package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/router/auth"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/sirupsen/logrus"
)

type UserLoginRes struct {
	Token string `json:"token,omitempty"`
}

func UserLogin(c *gin.Context) {
	userID, err := getUserLoginQueryParams(c)
	if err != nil {
		err = fmt.Errorf("bad query param: %w", err)
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	tokenString, err := auth.GenerateJWTString(userID)
	if err != nil {
		logrus.Errorf("failed to generate JWT string: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "failed to generate JWT string",
		})
		return
	}

	res := &UserLoginRes{
		Token: tokenString,
	}

	c.JSON(http.StatusOK, res)
}

func getUserLoginQueryParams(c *gin.Context) (string, error) {
	if c.Query(handler.HeaderUserID) == "" {
		return "", fmt.Errorf("user id is not given")

	}

	userID := c.Query(handler.HeaderUserID)
	return userID, nil
}
