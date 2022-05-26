package process

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manga-reader/manga-reader/backend/auth"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/reader"
)

func ProcessLoad(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		comicID := c.Param("comic_id")
		userToken := c.GetHeader("token")

		jwt, err := auth.DecodeJWT(userToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to Load process",
			})
		}

		proc, err := reader.GetReader(jwt.UserID, db).ProcessLoad(reader.Website_8comic, comicID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "failed to Load process",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"page": proc,
		})
	}
}
