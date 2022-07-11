package record

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/reader"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/sirupsen/logrus"
)

type RecordLoadRes struct {
	Volume string `json:"volume,omitempty"`
	Page   int    `json:"page,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

func RecordLoad(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID := getRecordLoadQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to decode JWT",
			})
		}

		var res RecordLoadRes
		res.Volume, res.Page, err = reader.GetReader(jwt.UserID, db).RecordLoad(reader.Website_8comic, comicID)
		if err != nil {
			logrus.Errorf("failed to Load record: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to Load record",
			})
		}
		c.JSON(http.StatusOK, res)
	}
}

func getRecordLoadQueryParams(c *gin.Context) string {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "comic id is not given"})
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID
}
