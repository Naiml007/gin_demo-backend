// routes/routes.go
package routes

import "github.com/gin-gonic/gin"

// SetupRoutes initializes all routes
func SetupRoutes(r *gin.Engine) {
	SetupHomeRoute(r)
	SetupPingRoute(r)
	SetupTestRoute(r)
	SetupUploadRoute(r)
	SetupMailRoute(r)
}
