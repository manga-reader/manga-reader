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

func RecordSave(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, vol, page := getRecordSaveQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to decode JWT",
			})
		}

		r := reader.GetReader(jwt.UserID, db)
		err = r.RecordSave(reader.Website_8comic, comicID, vol, page)
		if err != nil {
			logrus.Errorf("failed to save record: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to save record",
			})
		}
		c.String(http.StatusOK, handler.ResponseOK)
	}
}

func getRecordSaveQueryParams(c *gin.Context) (string, string, string) {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "comic id is not given"})
	}
	if c.Query(handler.HeaderVolume) == "" {
		logrus.Errorf("volume is not given")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "volume is not given"})
	}
	if c.Query(handler.HeaderPage) == "" {
		logrus.Errorf("page is not given")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "page is not given"})
	}
	comicID := c.Query(handler.HeaderComicID)
	vol := c.Query(handler.HeaderVolume)
	page := c.Query(handler.HeaderPage)
	return comicID, vol, page
}
