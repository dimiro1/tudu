package localstorage

import (
	"context"
	"sync"

	"github.com/dimiro1/tudu/internal/storage"
)

// InMemory static in memory implementation of tudu.Store
type InMemory struct {
	sync.RWMutex
	todos []storage.Todo
}

// NewInMemory returns a new InMemory instance.
func NewInMemory() *InMemory {
	return &InMemory{
		todos: []storage.Todo{
			{
				ID:        1,
				Title:     "Finish my homework",
				Completed: false,
			},
			{
				ID:        2,
				Title:     "Develop a new sample app",
				Completed: true,
			},
		},
	}
}

func (l *InMemory) All(context.Context) ([]storage.Todo, error) {
	l.RLock()
	defer l.RUnlock()
	return l.todos, nil
}

func (l *InMemory) Single(_ context.Context, id int) (storage.Todo, error) {
	l.RLock()
	defer l.RUnlock()

	for _, todo := range l.todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return storage.Todo{}, storage.ErrTodoNotFound
}
