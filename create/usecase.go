package todo

import (
	"context"
	"errors"
	"strings"
	"time"
)

// Sentinel errors for validation
var (
	ErrEmptyTitle     = errors.New("title cannot be empty")
	ErrInvalidDueDate = errors.New("due date must be in the future")
	ErrTodoNotValid   = errors.New("todo is not valid")
)

// Todo represents the structure of a todo item in the system.
type Todo struct {
	ID        string
	Title     string
	Completed bool
	DueDate   time.Time
}

// Validate checks if the todo's title is not empty and the due date is in the future.
// It returns a sentinel error corresponding to the validation failure.
func (t Todo) Validate() error {
	if strings.Trim(t.Title, " ") == "" {
		return ErrEmptyTitle
	}

	if !t.DueDate.After(time.Now()) {
		return ErrInvalidDueDate
	}

	return nil
}

// SaveTodo is a function type representing the save operation for a todo.
type SaveTodo func(ctx context.Context, todo Todo) error

// CreateTodoUseCase is a function type that defines the use case to create a todo.
type CreateTodoUseCase func(ctx context.Context, title string, dueDate time.Time) (Todo, error)

// NewCreateTodoUseCase returns a function that represents the todo creation use case.
// It takes a SaveTodo function as a parameter which will be used to persist the todo.
func NewCreateTodoUseCase(saveTodo SaveTodo) CreateTodoUseCase {
	return func(ctx context.Context, title string, dueDate time.Time) (Todo, error) {
		// Create todo item
		todo := Todo{
			Title:   title,
			DueDate: dueDate,
		}

		// Validate the todo
		if err := todo.Validate(); err != nil {
			return Todo{}, err
		}

		// Persist the todo using the provided saveTodo function
		if err := saveTodo(ctx, todo); err != nil {
			return Todo{}, err
		}

		return todo, nil
	}
}
