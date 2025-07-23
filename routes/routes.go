package routes

import (
	"todo/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	// Define the actual CRUD routes
	rg.GET("/", handlers.GetAll)
	rg.POST("/", handlers.AddTodo)
	rg.PUT("/:id", handlers.UpdateTodo)
	rg.DELETE("/:id", handlers.DeleteTodo)

	// Optional: Handle CORS preflight for all methods
	rg.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Origin")
		c.Status(204)
	})
}
