package main

import (
	"api/api/controllers"
	"api/api/models"
	repositories "api/api/repositories"
	"api/api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repositories.TodoMemoryDB{Todos: []models.Todo{}}
	inMemoryRepository := repositories.NewInMemoryTodoRepository(db)

	todoService := services.NewTodoService(inMemoryRepository)
	todoController := controllers.NewTodoController(*todoService)

	router := gin.Default()
	router.GET("/todos", todoController.Get)
	router.POST("/todo", todoController.Post)

	router.Run("localhost:3000")
}
