package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router/handler/health"
	"github.com/manga-reader/manga-reader/backend/router/handler/record"
	"github.com/manga-reader/manga-reader/backend/router/handler/user"
)

type Params struct {
	Database *database.Database
}

// SetupRouter Create a new router object
func SetupRouter(params *Params) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery())

	corsMiddleware := func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "User-Agent, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", http.MethodPost)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
	r.Use(corsMiddleware)

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
		userRoute.GET("/favorite", user.UserGetFavorite(params.Database))
		userRoute.POST("/favorite", user.UserAddFavorite(params.Database))
		userRoute.DELETE("/favorite", user.UserDelFavorite(params.Database))
		userRoute.GET("/history", user.UserGetHistory(params.Database))
	}

	recordRoute := r.Group("/record").Use(middlewareCheckJWTToken)
	{
		recordRoute.GET("/save", record.RecordSave(params.Database))
		recordRoute.GET("/load", record.RecordLoad(params.Database))
	}

	return r
}

func middlewareCheckJWTToken(c *gin.Context) {
	if !checkJWTToken() {
		c.Abort()
		c.JSON(200, gin.H{"msg": "jwt token is not valid"})
	}
	c.Next()
}

func checkJWTToken() bool {
	return true
}
