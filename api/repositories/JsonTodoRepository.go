package repositories

import (
	"api/api/models"
	"api/internal/database"
)

type JSONDBTodoRepository struct {
	db    database.DB[models.Todo]
	table string
}

func NewJSONDBTodoRepository(db database.DB[models.Todo]) *JSONDBTodoRepository {
	return &JSONDBTodoRepository{db: db, table: "todos"}
}

func (todoRepository *JSONDBTodoRepository) Find() ([]models.Todo, error) {
	return todoRepository.db.Get(todoRepository.table), nil
}

func (todoRepository *JSONDBTodoRepository) Create(todo models.Todo) (models.Todo, error) {
	todoRepository.db.Record(todoRepository.table, &todo)
	return todo, nil
}
