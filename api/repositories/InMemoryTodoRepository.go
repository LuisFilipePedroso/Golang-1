package repositories

import "api/api/models"

type TodoMemoryDB struct {
	Todos []models.Todo
}

type InMemoryTodoRepository struct {
	db TodoMemoryDB
}

func NewInMemoryTodoRepository(db TodoMemoryDB) *InMemoryTodoRepository {
	return &InMemoryTodoRepository{db: db}
}

func (todoRepository *InMemoryTodoRepository) Find() ([]models.Todo, error) {
	return todoRepository.db.Todos, nil
}

func (todoRepository *InMemoryTodoRepository) Create(todo models.Todo) (models.Todo, error) {
	todoRepository.db.Todos = append(todoRepository.db.Todos, todo)
	return todo, nil
}
