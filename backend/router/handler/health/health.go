package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// PerformHealthCheck Response pong for checking server status
func PerformHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "pong")

	logrus.Debug("pong")
}
