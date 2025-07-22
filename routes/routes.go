package routes

import (
	"todo/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", handlers.GetAll)
	rg.POST("/", handlers.AddTodo)
	rg.PUT("/:id", handlers.UpdateTodo)
	rg.DELETE("/:id", handlers.DeleteTodo)
}
