package record

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
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
		comicID := getRecordLoadQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to decode JWT",
			})
		}

		var res RecordLoadRes
		res.Volume, res.Page, err = u.RecordLoad(usecases.Website_8comic, jwt.UserID, comicID)
		if err != nil {
			logrus.Errorf("failed to Load record: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to Load record",
			})
		}
		c.JSON(http.StatusOK, res)
	}
}

func getRecordLoadQueryParams(c *gin.Context) string {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "comic id is not given"})
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID
}
