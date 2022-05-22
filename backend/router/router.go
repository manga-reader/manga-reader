package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router/handler/health"
	"github.com/manga-reader/manga-reader/backend/router/handler/process"
)

type Params struct {
	Database *database.Database
}

type Options struct {
}

// SetupRouter Create a new router object
func SetupRouter(params *Params, opts *Options) *gin.Engine {
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

	healthRoute := r.Group("/health")
	{
		healthRoute.GET("/ping", health.HealthPing)
	}

	processRoute := r.Group("/process")
	{
		processRoute.GET("/save/:vol/:page", process.ProcessSave)
	}

	return r
}
