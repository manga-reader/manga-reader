package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/manga-reader/manga-reader/backend/router/handler/health"
	"github.com/manga-reader/manga-reader/backend/router/handler/record"
	"github.com/manga-reader/manga-reader/backend/router/handler/user"
	"github.com/manga-reader/manga-reader/backend/usecases"
)

type Params struct {
	Usecase *usecases.Usecase
}

// SetupRouter Create a new router object
func SetupRouter(params *Params) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "User-Agent, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", http.MethodPost)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// show the called path
	r.Use(func(c *gin.Context) {
		logrus.Info("calling ", c.Request.URL)
		c.Next()
	})

	healthRoute := r.Group("/health")
	{
		healthRoute.GET("/ping", health.HealthPing)
	}

	userRoute := r.Group("/user")
	{
		userRoute.GET("/login", user.UserLogin)

		userRoute.GET("/favorite", user.UserGetFavorite(params.Usecase))
		userRoute.POST("/favorite", user.UserAddFavorite(params.Usecase))
		userRoute.DELETE("/favorite", user.UserDelFavorite(params.Usecase))

		userRoute.GET("/history", user.UserGetHistory(params.Usecase))
	}

	recordRoute := r.Group("/record")
	{
		recordRoute.GET("/new", record.RecordNew(params.Usecase))
		recordRoute.GET("/save", record.RecordSave(params.Usecase))
		recordRoute.GET("/load", record.RecordLoad(params.Usecase))
	}

	return r
}
