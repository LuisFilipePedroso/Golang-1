package controllers

import (
	"api/api/models"
	"api/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (controller *TodoController) Get(ctx *gin.Context) {
	todos, _ := controller.service.Find()
	ctx.IndentedJSON(http.StatusOK, todos)

	return
}

func (controller *TodoController) Post(ctx *gin.Context) {
	var bindedTodo models.Todo

	// Call BindJSON to bind the received JSON to
	// todo.
	if err := ctx.BindJSON(&bindedTodo); err != nil {
		return
	}

	todo, _ := controller.service.Create(bindedTodo.Title)
	ctx.IndentedJSON(http.StatusCreated, todo)
}
