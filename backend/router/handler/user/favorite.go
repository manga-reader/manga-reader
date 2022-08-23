package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/router/handler"
	"github.com/manga-reader/manga-reader/backend/usecases"
	"github.com/sirupsen/logrus"
)

func UserGetFavorite(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to decode JWT",
			})
		}

		list, err := u.GetFavorites(jwt.UserID, 0, 0)
		if err != nil {
			logrus.Errorf("failed to get user favorite: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to get user favorite",
			})
		}

		c.JSON(http.StatusOK, list)
	}
}

func UserAddFavorite(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID := geUserAddFavoriteQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to decode JWT",
			})
		}

		err = u.AddFavorite(jwt.UserID, comicID)
		if err != nil {
			logrus.Errorf("failed to add user favorite: %v: %v", comicID, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to add user favorite",
			})
		}
		c.JSON(http.StatusOK, gin.H{"msg": handler.ResponseOK})
	}
}

func UserDelFavorite(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID := geUserDelFavoriteQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v: %v", comicID, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to decode JWT",
			})
		}

		err = u.DelFavorite(jwt.UserID, comicID)
		if err != nil {
			logrus.Errorf("failed to delete user favorite: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to delete user favorite",
			})
		}
		c.JSON(http.StatusOK, gin.H{"msg": handler.ResponseOK})
	}
}

func geUserAddFavoriteQueryParams(c *gin.Context) string {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "comic id is not given"})
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID
}

func geUserDelFavoriteQueryParams(c *gin.Context) string {
	if c.Query(handler.HeaderComicID) == "" {
		logrus.Errorf("comic id is not given")
		c.JSON(http.StatusBadRequest, gin.H{"err": "comic id is not given"})
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID
}
