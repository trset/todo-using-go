package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"todo/routes" // adjust if using a different import path
)

func StartServer() {
	r := gin.Default()

	// ✅ Correct CORS setup to allow frontend hosted at GitHub Pages
	config := cors.Config{
		AllowOrigins:     []string{"https://trset.github.io"}, // ✅ GitHub Pages frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	// ✅ Serve static frontend files if needed (can skip on Render if only serving API)
	// r.Static("/", "./")

	// ✅ Register API routes under /api (this is important to match the frontend fetch)
	api := r.Group("/api")
	{
		routes.RegisterRoutes(api)
	}

	// ✅ Optional fallback route (can serve 404 or homepage)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	// ✅ Start the server
	r.Run() // defaults to ":8080" or uses PORT env var on Render
}
