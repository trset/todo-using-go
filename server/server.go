package server

import (
	"todo/routes" // or "todo-using-go/routes" if that's your actual module path

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// Apply proper CORS middleware
	r.Use(cors.Default())

	// Serve static files
	r.StaticFile("/", "./frontend/index.html")
	r.Static("/static", "./frontend")

	r.StaticFile("/manifest.json", "./frontend/manifest.json")
	r.StaticFile("/sw.js", "./frontend/sw.js")
	r.StaticFile("/icon-192.png", "./frontend/icon-192.png")
	r.StaticFile("/icon-512.png", "./frontend/icon-512.png")

	// API routes
	api := r.Group("/api")
	routes.RegisterRoutes(api)

	// SPA fallback
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	r.Run()
}
