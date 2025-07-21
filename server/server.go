package server

import (
	"todo/routes" // adjust if your module path is different

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins for dev (you can restrict in prod)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	
	r.StaticFile("/", "./index.html")
	r.StaticFile("/manifest.json", "./manifest.json")
	r.StaticFile("/service-worker.js", "./service-worker.js")
	r.StaticFile("/icon-192.png", "./icon-192.png")
	r.StaticFile("/icon-512.png", "./icon-512.png")


	// Register routes
	routes.RegisterRoutes(r)

	// Start server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
