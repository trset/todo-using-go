package server

import (
	"todo/routes" // adjust if your module path is different

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// ✅ CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // You can restrict this later in production
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Serve static files and index.html
	r.StaticFile("/", "./index.html")
	r.StaticFile("/index.html", "./index.html")
	r.StaticFile("/manifest.json", "./manifest.json")
	r.StaticFile("/sw.js", "./sw.js")
	r.StaticFile("/icon-192.png", "./icon-192.png")
	r.StaticFile("/icon-512.png", "./icon-512.png")

	// Global OPTIONS handler for CORS preflight
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	// Register API routes
	api := r.Group("/api")
	routes.RegisterRoutes(api)

	// ✅ Start the server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
