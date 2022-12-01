package services

import "api/api/models"

type TodoService struct {
	Repository models.TodoRepository
}

func NewTodoService(repository models.TodoRepository) *TodoService {
	return &TodoService{Repository: repository}
}

func (t *TodoService) Find() ([]models.Todo, error) {
	todos, err := t.Repository.Find()

	if err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}

func (t *TodoService) Create(title string) (models.Todo, error) {
	newTodo := models.NewTodo()
	newTodo.Title = title

	todo, err := t.Repository.Create(*newTodo)

	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}
