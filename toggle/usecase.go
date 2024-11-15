package toggle

import (
	"context"
	"errors"
)

var ErrNoID = errors.New("todo ID cannot be empty")

// ToggleTodosUseCase is a function type that defines the use case to toggle the completion status of a todo.
type ToggleTodosUseCase func(ctx context.Context, id string) error

// ToggleTodo represents the operation to toggle a todo's completion status in the data store.
type ToggleTodo func(ctx context.Context, id string) error

// NewToggleTodosUseCase returns a function that represents the todo toggle use case.
// It takes a ToggleTodo function as a parameter which will be used to toggle the todo status.
func NewToggleTodosUseCase(toggleTodo ToggleTodo) ToggleTodosUseCase {
	return func(ctx context.Context, id string) error {
		if id == "" {
			return ErrNoID
		}

		return toggleTodo(ctx, id)
	}
}
