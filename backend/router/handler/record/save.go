package record

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/sirupsen/logrus"
)

func RecordSave(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, vol, page := getRecordSaveQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to decode JWT",
			})
		}

		err = u.RecordSave(usecases.Website_8comic, jwt.UserID, comicID, vol, page)
		if err != nil {
			logrus.Errorf("failed to save record: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to save record",
			})
		}
		c.String(http.StatusOK, handler.ResponseOK)
	}
}

func getRecordSaveQueryParams(c *gin.Context) (string, string, int) {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "comic id is not given"})
	}
	if c.Query(handler.HeaderVolume) == "" {
		logrus.Errorf("volume is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "volume is not given"})
	}
	if c.Query(handler.HeaderPage) == "" {
		logrus.Errorf("page is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "page is not given"})
	}
	comicID := c.Query(handler.HeaderComicID)
	vol := c.Query(handler.HeaderVolume)
	pageRaw := c.Query(handler.HeaderPage)
	page, err := strconv.Atoi(pageRaw)
	if err != nil {
		logrus.Errorf("page: '%s' is not a number", pageRaw)
		c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("page: '%s' is not a number", pageRaw)})
	}
	return comicID, vol, page
}
