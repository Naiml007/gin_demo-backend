// routes/ping.go
package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupPingRoute initializes the ping route
func SetupPingRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		RenderCommonTemplate(c, "Main website", "Ping Route", "pong", "")
	})
}
