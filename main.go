package main

import (
	"api/api/controllers"
	"api/api/models"
	repositories "api/api/repositories"
	"api/api/services"
	"api/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := repositories.TodoMemoryDB{Todos: []models.Todo{}}
	// inMemoryRepository := repositories.NewInMemoryTodoRepository(db)

	db := database.DB[models.Todo]{}
	jsonDBRepository := repositories.NewJSONDBTodoRepository(db)

	todoService := services.NewTodoService(jsonDBRepository)
	todoController := controllers.NewTodoController(*todoService)

	router := gin.Default()
	router.GET("/todos", todoController.Get)
	router.POST("/todo", todoController.Post)

	router.Run("localhost:3000")
}
