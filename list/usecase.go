package list

import (
	"context"
	"errors"
	"time"
)

var (
	ErrPageNumberInvalid = errors.New("page number must be greater than 0")
	ErrPageSizeInvalid   = errors.New("page size must be greater than 0")
	ErrNoTodosFound      = errors.New("no todos found")
)

// Todo represents the structure of a todo item in the system.
type Todo struct {
	ID        string
	Title     string
	Completed bool
	DueDate   time.Time
}

// PageOptions represents the options for pagination.
type PageOptions struct {
	Number int
	Size   int
}

// Validate checks if the page options are valid.
func (p *PageOptions) Validate() error {
	if p.Number < 1 {
		return ErrPageNumberInvalid
	}

	if p.Size < 1 {
		return ErrPageSizeInvalid
	}

	return nil
}

// Pages represents the paginated todos and total count.
type Pages struct {
	Todos []Todo
	Page  int
	Size  int
	Total int
}

// ListTodosUseCase is a function type that defines the use case to list todos with pagination.
type ListTodosUseCase func(ctx context.Context, filterByCompleted bool, opts *PageOptions) (Pages, error)

// GetTodosWithPagination is a function type representing the operation to fetch todos from the data store with pagination.
type GetTodosWithPagination func(ctx context.Context, filterByCompleted bool, opts PageOptions) ([]Todo, error)

// NewListTodosUseCase returns a function that represents the todo listing use case with pagination.
// It takes a GetTodosWithPagination function as a parameter which will be used to fetch todos with pagination.
func NewListTodosUseCase(getTodos GetTodosWithPagination) ListTodosUseCase {
	return func(ctx context.Context, filterByCompleted bool, opts *PageOptions) (Pages, error) {
		if opts != nil {
			// page options where provided, validate them
			if err := opts.Validate(); err != nil {
				return Pages{}, err
			}
		} else {
			// no page options provided, use default values
			opts = &PageOptions{
				Number: 1,
				Size:   10,
			}
		}

		todos, err := getTodos(ctx, filterByCompleted, *opts)
		if err != nil {
			return Pages{}, err
		}

		// Construct the Pages output
		return Pages{
			Todos: todos,
			Page:  opts.Number,
			Size:  opts.Size,
			Total: len(todos),
		}, nil
	}
}
