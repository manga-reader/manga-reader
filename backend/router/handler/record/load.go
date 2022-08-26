package record

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/router/auth"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/sirupsen/logrus"
)

type RecordLoadRes struct {
	Volume string `json:"volume,omitempty"`
	Page   int    `json:"page,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

func RecordLoad(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, err := getRecordLoadQueryParams(c)
		if err != nil {
			err = fmt.Errorf("bad query param: %w", err)
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "failed to decode JWT",
			})
			return
		}

		var res RecordLoadRes
		res.Volume, res.Page, err = u.RecordLoad(usecases.Website_8comic, jwt.UserID, comicID)
		if err != nil {
			logrus.Errorf("failed to Load record: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to Load record",
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func getRecordLoadQueryParams(c *gin.Context) (string, error) {
	if c.Query(handler.HeaderComicID) == "" {
		return "", fmt.Errorf("comic id is not given")
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID, nil
}
