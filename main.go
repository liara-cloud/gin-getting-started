package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files (CSS, JS, images)
	r.Static("/css", "public/css")
	r.Static("/js", "public/js")
	r.Static("/images", "public/images")

	// Serve the HTML file at the root
	r.StaticFile("/", "public/index.html")

	r.Run(":8080") // Start the server on port 8080
}
