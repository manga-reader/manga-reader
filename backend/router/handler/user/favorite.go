package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/router/auth"
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
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "failed to decode JWT",
			})
			return
		}

		list, err := u.GetFavorites(jwt.UserID, 0, 0)
		if err != nil {
			logrus.Errorf("failed to get user favorite: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to get user favorite",
			})
			return
		}

		c.JSON(http.StatusOK, list)
	}
}

func UserAddFavorite(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, err := geUserAddFavoriteQueryParams(c)
		if err != nil {
			err = fmt.Errorf("bad query param: %w", err)
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
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

		err = u.AddFavorite(jwt.UserID, comicID)
		if err != nil {
			logrus.Errorf("failed to add user favorite: %v: %v", comicID, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to add user favorite",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": handler.ResponseOK})
	}
}

func UserDelFavorite(u *usecases.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, err := geUserDelFavoriteQueryParams(c)
		if err != nil {
			err = fmt.Errorf("bad query param: %w", err)
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			logrus.Errorf("failed to decode JWT: %v: %v", comicID, err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "failed to decode JWT",
			})
			return
		}

		err = u.DelFavorite(jwt.UserID, comicID)
		if err != nil {
			logrus.Errorf("failed to delete user favorite: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "failed to delete user favorite",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": handler.ResponseOK})
	}
}

func geUserAddFavoriteQueryParams(c *gin.Context) (string, error) {
	if c.Query(handler.HeaderComicID) == "" {
		return "", fmt.Errorf("comic id is not given")
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID, nil
}

func geUserDelFavoriteQueryParams(c *gin.Context) (string, error) {
	if c.Query(handler.HeaderComicID) == "" {
		return "", fmt.Errorf("comic id is not given")
	}

	comicID := c.Query(handler.HeaderComicID)
	return comicID, nil
}
