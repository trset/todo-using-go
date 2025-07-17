package routes

import (
	"todo/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/", handlers.AddTodo)
	r.GET("/", handlers.GetAll)
	r.GET("/:id", handlers.GetById)
	r.PUT("/:id", handlers.UpdateTodo)
	r.DELETE("/:id", handlers.DeleteTodo)
}
