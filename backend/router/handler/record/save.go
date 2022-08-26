package record

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/router/auth"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/sirupsen/logrus"
)

func RecordSave(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, vol, page, err := getRecordSaveQueryParams(c)
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

		err = u.RecordSave(usecases.Website_8comic, jwt.UserID, comicID, vol, page)
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

func getRecordSaveQueryParams(c *gin.Context) (string, string, int, error) {
	if c.Query(handler.HeaderComicID) == "" {
		return "", "", 0, fmt.Errorf("comic id is not given")
	}
	if c.Query(handler.HeaderVolume) == "" {
		return "", "", 0, fmt.Errorf("volume is not given")
	}
	if c.Query(handler.HeaderPage) == "" {
		return "", "", 0, fmt.Errorf("page is not given")
	}
	comicID := c.Query(handler.HeaderComicID)
	vol := c.Query(handler.HeaderVolume)
	pageRaw := c.Query(handler.HeaderPage)
	page, err := strconv.Atoi(pageRaw)
	if err != nil {
		logrus.Errorf("page: '%s' is not a number", pageRaw)
		c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("page: '%s' is not a number", pageRaw)})
	}
	return comicID, vol, page, nil
}
