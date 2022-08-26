package record

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/crawler"
	"github.com/manga-reader/manga-reader/backend/router/auth"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/sirupsen/logrus"
)

func RecordNew(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, vol := getRecordNewQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to decode JWT",
			})
		}

		name, latestVol, updatedAt, err := crawler.GetComicInfo(comicID)
		if err != nil {
			logrus.Errorf("can't get comic info of %s in RecordNew(): %s", comicID, err)
			c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("can't get comic info of %s in RecordNew(): %s", comicID, err)})
		}

		err = u.AddComic(comicID, name, latestVol, updatedAt)
		if err != nil {
			logrus.Errorf("can't add new comic of %s in RecordNew(): %s", comicID, err)
			c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("can't add new comic of %s in RecordNew(): %s", comicID, err)})
		}

		err = u.RecordSave(usecases.Website_8comic, jwt.UserID, comicID, vol, 0)
		if err != nil {
			logrus.Errorf("failed to save record: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to save record",
			})
		}
		c.String(http.StatusOK, handler.ResponseOK)
	}
}

func getRecordNewQueryParams(c *gin.Context) (string, string) {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "comic id is not given"})
	}
	if c.Query(handler.HeaderVolume) == "" {
		logrus.Errorf("volume is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "volume is not given"})
	}

	comicID := c.Query(handler.HeaderComicID)
	vol := c.Query(handler.HeaderVolume)

	return comicID, vol
}
