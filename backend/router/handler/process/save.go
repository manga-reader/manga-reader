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
		comicID := c.Param("comic_id")
		vol := c.Param("vol")
		page := c.Param("page")
		userToken := c.GetHeader("token")

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to save process",
			})
		}

		err = reader.GetReader(jwt.UserID, db).ProcessSave(reader.Website_8comic, comicID, vol, page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to save process",
			})
		}
		c.String(http.StatusOK, handler.ResponseOK)
	}
}
