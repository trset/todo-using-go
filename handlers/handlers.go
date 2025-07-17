package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	tl = []Todo{}
)

func AddTodo(c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	newTodo := Todo{
		ID:    uuid.New().String(),
		Title: input.Title,
		Done:  false,
	}

	tl = append(tl, newTodo)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo added successfully",
		"todo":    newTodo,
	})
}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"todos": tl,
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id") // get path param

	for _, todo := range tl {
		if todo.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"todo": todo,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Todo not found",
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Title string `json:"title"`
		Done  *bool  `json:"done"` // pointer to detect if provided
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, todo := range tl {
		if todo.ID == id {
			if input.Title != "" {
				tl[i].Title = input.Title
			}
			if input.Done != nil {
				tl[i].Done = *input.Done
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Todo updated",
				"todo":    tl[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range tl {
		if todo.ID == id {
			tl = append(tl[:i], tl[i+1:]...) // remove the item
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
