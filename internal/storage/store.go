package storage

import (
	"context"
	"errors"
)

var ErrTodoNotFound = errors.New("todo not found")

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Store interface {
	// All returns all todos.
	All(ctx context.Context) ([]Todo, error)

	// Single returns a single todo item.
	// Returns ErrTodoNotFound if no todo was found.
	Single(ctx context.Context, id int) (Todo, error)
}
