// routes/home.go
package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupHomeRoute initializes the home route
func SetupHomeRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		// Render the common template with home route content
		RenderCommonTemplate(c, "Main website", "Home Route", "This is the content for the home route.", "")
	})
}
