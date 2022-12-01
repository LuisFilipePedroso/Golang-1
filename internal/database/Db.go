package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	. "api/api/models"

	"github.com/google/uuid"
)

var FILENAME = "db.json"

func Get(table string) []Todo {
	data, err := os.Open(FILENAME)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	jsonData, error := ioutil.ReadAll(data)

	if error != nil {
		panic("Error on read table")
	}

	result := map[string][]Todo{}

	json.Unmarshal([]byte(jsonData), &result)

	return result[table]
}

func Record(table string, todo *Todo) {
	db := Get(table)

	id := uuid.New()

	todo.ID = id.String()

	appendedData := append(db, *todo)

	result := map[string][]Todo{
		table: appendedData,
	}

	file, _ := json.MarshalIndent(result, "", " ")

	_ = ioutil.WriteFile(FILENAME, file, 0644)
}
