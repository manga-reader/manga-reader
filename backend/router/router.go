package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router/handler/health"
	"github.com/manga-reader/manga-reader/backend/router/handler/process"
	reader "github.com/manga-reader/manga-reader/backend/router/handler/user"
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

	healthRoute := r.Group("/health")
	{
		healthRoute.GET("/ping", health.HealthPing)
	}

	userRoute := r.Group("/user")
	{
		userRoute.GET("/login", reader.UserLogin)
	}

	processRoute := r.Group("/process").Use(middlewareCheckJWTToken)
	{
		processRoute.GET("/save/:comic_id/:vol/:page", process.ProcessSave(params.Database))
		processRoute.GET("/load/:comic_id", process.ProcessLoad(params.Database))
	}

	return r
}

func middlewareCheckJWTToken(c *gin.Context) {
	if checkJWTToken() {
		c.Next()
	}

	c.Abort()
	c.JSON(200, gin.H{"msg": "jwt token is not valid"})
}

func checkJWTToken() bool {
	return true
}
