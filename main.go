package main

import (
	"github.com/gin-gonic/gin"
	"todo/routes"
)

func main() {
	r := gin.Default()

	// Serve static frontend files
	r.StaticFile("/", "./index.html") // This serves index.html at root
	r.Static("/static", ".")          // Serve all other static files like manifest, icons, sw.js

	// Register API routes
	routes.RegisterRoutes(r.Group("/api"))

	r.Run() // default port 8080
}
