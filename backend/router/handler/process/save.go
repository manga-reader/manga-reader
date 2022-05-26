package process

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/reader"
	"github.com/manga-reader/manga-reader/backend/router/handler"
)

func ProcessSave(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID, vol, page := getProcessSaveQueryParams(c)
		userToken := c.GetHeader(handler.HeaderJWTToken)

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to save process",
			})
		}

		r := reader.GetReader(jwt.UserID, db)
		err = r.ProcessSave(reader.Website_8comic, comicID, vol, page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to save process",
			})
		}
		c.String(http.StatusOK, handler.ResponseOK)
	}
}

func getProcessSaveQueryParams(c *gin.Context) (string, string, string) {
	if c.Query(handler.HeaderComicID) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "comic id is not given"})
	}
	if c.Query(handler.HeaderVolume) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "volume is not given"})
	}
	if c.Query(handler.HeaderPage) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "page is not given"})
	}
	comicID := c.Query(handler.HeaderComicID)
	vol := c.Query(handler.HeaderVolume)
	page := c.Query(handler.HeaderPage)

	return comicID, vol, page
}
