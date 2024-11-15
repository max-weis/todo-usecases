package remove

import (
	"context"
	"errors"
)

var ErrNoID = errors.New("todo ID cannot be empty")

// RemoveTodoUseCase is a function type that defines the use case to delete a todo.
type RemoveTodoUseCase func(ctx context.Context, id string) error

// RemoveTodo represents the operation to delete a todo from the data store.
type RemoveTodo func(ctx context.Context, id string) error

// NewRemoveTodoUseCase returns a function that represents the todo deletion use case.
// It takes a RemoveTodo function as a parameter which will be used to delete the todo.
func NewRemoveTodoUseCase(deleteTodo RemoveTodo) RemoveTodoUseCase {
	return func(ctx context.Context, id string) error {
		// Input validation
		if id == "" {
			return ErrNoID
		}

		// Remove the todo
		if err := deleteTodo(ctx, id); err != nil {
			return err
		}

		return nil
	}
}
