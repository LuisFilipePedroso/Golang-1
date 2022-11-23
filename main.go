package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var FILENAME = "db.json"

type Todo struct {
	ID         string `json:"id"`
	title      string `json:"title"`
	isComplete string `json:"isComplete"`
}

func getTableData(table string) []Todo {
	data, err := os.Open(FILENAME)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	jsonData, _ := ioutil.ReadAll(data)

	result := map[string][]Todo{}

	json.Unmarshal([]byte(jsonData), &result)

	return result[table]
}

func getTodos(c *gin.Context) {
	todos := getTableData("todos")

	c.IndentedJSON(http.StatusOK, todos)
}

// func createAlbum(c *gin.Context) {
// 	var newAlbum Album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	// router.POST("/album", createAlbum)

	router.Run("localhost:3000")

}
