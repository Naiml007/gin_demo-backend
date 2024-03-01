// routes/test.go
package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupTestRoute initializes the test route
func SetupTestRoute(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		// Render common template with JavaScript redirect to "http://www.google.com/" in a new tab
		RenderCommonTemplate(c, "Main website", "Test Route", "", `
            <script>
                window.open("http://www.google.com/", "_blank");
            </script>
        `)
	})
}
