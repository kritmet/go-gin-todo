package repository

import (
	"encoding/json"
	"os"

	"github.com/kritmet/go-gin-todo/domain"
)

// JsonRepository json repository
type TodoRepository struct{}

// NewTodoRepository is a function for create todo repository
func NewTodoRepository() domain.TodoRepository {
	return &TodoRepository{}
}

const (
	todoJson string = "assets/todo.json"
)

// WriteJSON write todo file
func (r *TodoRepository) WriteTodoJSON(entities []*domain.Todo) error {
	jsonData, err := json.MarshalIndent(entities, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(todoJson, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ReadTodoJSON read todo file
func (r *TodoRepository) ReadTodoJSON() ([]*domain.Todo, error) {
	jsonData, err := os.ReadFile(todoJson)
	if err != nil {
		return nil, err
	}

	entities := []*domain.Todo{}
	err = json.Unmarshal(jsonData, &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}
