package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route handler for the GET /todos endpoint
	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get all todos",
		})
	})

	// Define a route handler for the GET /todos/:id endpoint
	router.GET("/todos/:id", func(c *gin.Context) {
		todoID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Get todo by ID",
			"id":      todoID,
		})
	})

	// Define a route handler for the POST /todos/add endpoint
	router.POST("/todos/add", func(c *gin.Context) {
		var todo struct {
			Title string `json:"title"`
			Done  bool   `json:"done"`
		}
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "Add a new todo",
			"title":   todo.Title,
			"done":    todo.Done,
		})
	})

	// Run the server
	router.Run(":8080")
}
