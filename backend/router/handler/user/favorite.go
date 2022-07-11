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

func UserFavorite(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to decode JWT",
			})
		}

		list, err := reader.GetReader(jwt.UserID, db).GetFavoriteList()
		if err != nil {
			logrus.Errorf("failed to get user favorite: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to get user favorite",
			})
		}
		c.JSON(http.StatusOK, list)
	}
}
