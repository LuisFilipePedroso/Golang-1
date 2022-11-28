package models

type Todo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete string `json:"isComplete"`
}
