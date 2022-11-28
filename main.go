package main

import (
	TodoController "api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", TodoController.GetTodos)
	router.POST("/todo", TodoController.CreateTodo)

	router.Run("localhost:3000")

}
