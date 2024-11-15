package search_test

import (
	"context"
	"errors"
	"github.com/max-weis/todo-usecases/search"
	"strings"
	"testing"
)

func TestSearchTodosUseCase(t *testing.T) {
	tests := []struct {
		name         string
		query        string
		expectedIDs  []string
		searchErr    error
		expectSearch bool
	}{
		{"Empty Query", "", []string{}, search.ErrNoSearchQuery, false},
		{"Valid Query", "buy", []string{"todo-1"}, nil, true},
		{"No Results", "non-existent", []string{}, nil, true},
		{"Search Error", "error", []string{}, errors.New("search error"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchFunc := func(ctx context.Context, query string) ([]search.Todo, error) {
				if query == "error" {
					return nil, tt.searchErr
				}
				// Mock todos based on the query for testing
				todos := []search.Todo{}
				if query == "" || strings.Contains("buy", query) {
					todos = append(todos, search.Todo{ID: "todo-1", Title: "Buy Milk"})
				}
				if query == "" {
					todos = append(todos, search.Todo{ID: "todo-2", Title: "Walk Dog"})
				}
				return todos, nil
			}
			useCase := search.NewSearchTodosUseCase(searchFunc)
			todos, err := useCase(context.Background(), tt.query)

			if tt.expectSearch && err != nil {
				t.Errorf("SearchTodosUseCase() error = %v, want nil", err)
				return
			}

			if !tt.expectSearch && err == nil {
				t.Errorf("SearchTodosUseCase() should have returned an error, got nil")
				return
			}

			if len(todos) != len(tt.expectedIDs) {
				t.Errorf("Unexpected number of todos: got %d, want %d", len(todos), len(tt.expectedIDs))
			}

			for i, todo := range todos {
				if todo.ID != tt.expectedIDs[i] {
					t.Errorf("Todo ID mismatch: got %s, want %s", todo.ID, tt.expectedIDs[i])
				}
			}
		})
	}
}
