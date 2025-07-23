package server

import (
	"todo/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func StartServer() {
	r := gin.Default()

	// ✅ CORS config
	config := cors.Config{
		AllowOrigins:     []string{"https://trset.github.io"}, // ✅ your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	// Serve static files
	r.Static("/", "./")
	
	// API routes
	api := r.Group("/api")
	{
		routes.RegisterRoutes(api)
	}

	// Fallback route
	r.NoRoute(func(c *gin.Context) {
		c.File("./index.html")
	})

	r.Run() // starts server on :8080
}
