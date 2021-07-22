package monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HeartbeatHandler - heartbeat endpoint
func HeartbeatHandler(router *gin.Engine) {
	router.GET("/heartbeat", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
