package search

import (
	"context"
	"errors"
	"strings"
	"time"
)

var ErrNoSearchQuery = errors.New("search query cannot be empty")

// Todo represents the structure of a todo item in the system.
type Todo struct {
	ID        string
	Title     string
	Completed bool
	DueDate   time.Time
}

// SearchTodosUseCase is a function type that defines the use case to search for todos.
type SearchTodosUseCase func(ctx context.Context, query string) ([]Todo, error)

// SearchTodos represents the operation to search todos in the data store.
type SearchTodos func(ctx context.Context, query string) ([]Todo, error)

// NewSearchTodosUseCase returns a function that represents the search todos use case.
// It takes a SearchTodos function as a parameter which will be used to search todos.
func NewSearchTodosUseCase(searchTodos SearchTodos) SearchTodosUseCase {
	return func(ctx context.Context, query string) ([]Todo, error) {
		// Trim the query to remove leading and trailing whitespace
		query = strings.TrimSpace(query)

		// Input validation
		if query == "" {
			return nil, ErrNoSearchQuery
		}

		// Execute the search with the provided query
		todos, err := searchTodos(ctx, query)
		if err != nil {
			return nil, err
		}

		return todos, nil
	}
}
