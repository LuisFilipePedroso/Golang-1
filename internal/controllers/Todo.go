package controllers

import (
	Db "api/internal/database"
	. "api/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var TABLENAME = "todos"

func GetTodos(c *gin.Context) {
	todos := Db.Get(TABLENAME)
	c.IndentedJSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo Todo

	// Call BindJSON to bind the received JSON to
	// todo.
	if err := c.BindJSON(&todo); err != nil {
		return
	}

	// Add the new todo to the slice.
	// albums = append(albums, newAlbum)
	Db.Record(TABLENAME, &todo)
	fmt.Println(todo)
}
