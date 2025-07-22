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

	// ✅ Serve static files (HTML/CSS/JS/icons)
	r.Static("/static", "./static")

	// ✅ Serve PWA-related files (e.g., manifest, service worker)
	r.StaticFile("/manifest.json", "./manifest.json")
	r.StaticFile("/service-worker.js", "./service-worker.js")
	r.StaticFile("/icon-192.png", "./icon-192.png")
	r.StaticFile("/icon-512.png", "./icon-512.png")

	// ✅ Serve index.html at root
	r.StaticFile("/", "./static/index.html")

	// ✅ Register routes under /api to avoid conflicts with /
	api := r.Group("/api")
	routes.RegisterRoutes(api)

	// ✅ Start the server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
