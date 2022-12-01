package models

import (
	uuid "github.com/google/uuid"
)

type TodoRepository interface {
	Find() ([]Todo, error)
	Create(todo Todo) (Todo, error)
}

type Todo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"isComplete"`
}

func NewTodo() *Todo {
	todo := Todo{
		ID:         uuid.New().String(),
		IsComplete: false,
	}

	return &todo
}
