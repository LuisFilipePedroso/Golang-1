package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var FILENAME = "db.json"

type DB[T any] struct{}

func NewDB[T any](db T) *DB[T] {
	return &DB[T]{}
}

func (d *DB[T]) Get(table string) []T {
	data, err := os.Open(FILENAME)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	jsonData, error := ioutil.ReadAll(data)

	if error != nil {
		panic("Error on read table")
	}

	result := map[string][]T{}

	json.Unmarshal([]byte(jsonData), &result)

	return result[table]
}

func (d *DB[T]) Record(table string, payload *T) {
	db := d.Get(table)

	appendedData := append(db, *payload)

	result := map[string][]T{
		table: appendedData,
	}

	file, _ := json.MarshalIndent(result, "", " ")

	_ = ioutil.WriteFile(FILENAME, file, 0644)
}
