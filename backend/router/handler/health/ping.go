package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// HealthPing Response pong for checking server status
func HealthPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")

	logrus.Debug("pong")
}
