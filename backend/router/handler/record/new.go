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
		comicID, vol, err := getRecordNewQueryParams(c)
		if err != nil {
			err = fmt.Errorf("bad query param: %w", err)
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
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

		name, latestVol, updatedAt, err := crawler.GetComicInfo(comicID)
		if err != nil {
			logrus.Errorf("can't get comic info of %s in RecordNew(): %s", comicID, err)
			c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("can't get comic info of %s in RecordNew(): %s", comicID, err)})
			return
		}

		err = u.AddComic(comicID, name, latestVol, updatedAt)
		if err != nil {
			logrus.Errorf("can't add new comic of %s in RecordNew(): %s", comicID, err)
			c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("can't add new comic of %s in RecordNew(): %s", comicID, err)})
			return
		}

		err = u.RecordSave(usecases.Website_8comic, jwt.UserID, comicID, vol, 0)
		if err != nil {
			logrus.Errorf("failed to save record: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to save record",
			})
			return
		}
		c.String(http.StatusOK, handler.ResponseOK)
	}
}

func getRecordNewQueryParams(c *gin.Context) (string, string, error) {
	if c.Query(handler.HeaderComicID) == "" {
		return "", "", fmt.Errorf("comic id is not given")
	}
	if c.Query(handler.HeaderVolume) == "" {
		return "", "", fmt.Errorf("volume is not given")
	}

	comicID := c.Query(handler.HeaderComicID)
	vol := c.Query(handler.HeaderVolume)

	return comicID, vol, nil
}
