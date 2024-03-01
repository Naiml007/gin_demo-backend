// main.go
package main

import (
	"fmt"
	"idkforNow/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	hello()
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Set up routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run() // listen and serve on 0.0.0.0:8080
}

func hello() {
	fmt.Println("Hello, world.")
}
