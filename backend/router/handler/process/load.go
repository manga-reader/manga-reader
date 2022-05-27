package process

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/reader"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/sirupsen/logrus"
)

type ProcessLoadRes struct {
	Volume string `json:"volume,omitempty"`
	Page   int    `json:"page,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

func ProcessLoad(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID := getProcessLoadQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to decode JWT",
			})
		}

		var res ProcessLoadRes
		res.Volume, res.Page, err = reader.GetReader(jwt.UserID, db).ProcessLoad(reader.Website_8comic, comicID)
		if err != nil {
			logrus.Errorf("failed to Load process: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to Load process",
			})
		}
		c.JSON(http.StatusOK, res)
	}
}

func getProcessLoadQueryParams(c *gin.Context) string {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "comic id is not given"})
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID
}
