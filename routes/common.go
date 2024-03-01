// routes/common.go
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RenderCommonTemplate renders the common HTML template with dynamic content
func RenderCommonTemplate(c *gin.Context, title, routeTitle, routeContent, additionalHTML string) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":          title,
		"routeTitle":     routeTitle,
		"routeContent":   routeContent,
		"additionalHTML": additionalHTML,
	})
}
