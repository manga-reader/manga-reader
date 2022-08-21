package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/reader"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/sirupsen/logrus"
)

func UserGetHistory(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to decode JWT",
			})
		}

		list, err := reader.GetReader(jwt.UserID, db).GetHistoryList(db)
		if err != nil {
			logrus.Errorf("failed to get user history: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to get user history",
			})
		}
		c.JSON(http.StatusOK, list)
	}
}