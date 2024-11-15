package list_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/max-weis/todo-usecases/list"
)

// Mock time variable to control due dates in tests
var now = time.Now()

// TestPageOptions_Validate tests the Validate method of PageOptions.
func TestPageOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    list.PageOptions
		wantErr error
	}{
		{"Valid Options", list.PageOptions{Number: 1, Size: 10}, nil},
		{"Invalid Page Number", list.PageOptions{Number: 0, Size: 10}, list.ErrPageNumberInvalid},
		{"Invalid Page Size", list.PageOptions{Number: 1, Size: 0}, list.ErrPageSizeInvalid},
		{"Invalid Both", list.PageOptions{Number: 0, Size: 0}, list.ErrPageNumberInvalid}, // Only one error is returned
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.opts.Validate(); !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestListTodosUseCase tests the returned ListTodosUseCase function.
func TestListTodosUseCase(t *testing.T) {
	dbError := errors.New("database error")
	tests := []struct {
		name          string
		filter        bool
		opts          *list.PageOptions
		getTodosErr   error
		expectedErr   error
		expectedTodos int
	}{
		{"Valid Options", false, &list.PageOptions{Number: 1, Size: 10}, nil, nil, 10},
		{"Invalid Page Number", false, &list.PageOptions{Number: 0, Size: 10}, nil, list.ErrPageNumberInvalid, 0},
		{"Invalid Page Size", false, &list.PageOptions{Number: 1, Size: 0}, nil, list.ErrPageSizeInvalid, 0},
		{"Default Options", false, nil, nil, nil, 10},
		{"GetTodos Error", false, &list.PageOptions{Number: 1, Size: 10}, dbError, dbError, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getTodos := func(ctx context.Context, filter bool, opts list.PageOptions) ([]list.Todo, error) {
				if tt.getTodosErr != nil {
					return nil, tt.getTodosErr
				}
				// For simplicity, let's assume we return a slice of Todos of length Size
				return make([]list.Todo, opts.Size), nil
			}
			useCase := list.NewListTodosUseCase(getTodos)
			pages, err := useCase(context.Background(), tt.filter, tt.opts)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("ListTodosUseCase() error = %v, wantErr %v", err, tt.expectedErr)
			}
			if tt.expectedErr == nil && len(pages.Todos) != tt.expectedTodos {
				t.Errorf("Unexpected number of todos: got %d, want %d", len(pages.Todos), tt.expectedTodos)
			}
		})
	}
}

// TestPages tests the construction of the Pages struct.
func TestPages(t *testing.T) {
	todos := []list.Todo{
		{ID: "1", Title: "Test Todo", Completed: false, DueDate: now.Add(24 * time.Hour)},
	}

	pages := list.Pages{
		Todos: todos,
		Page:  1,
		Size:  10,
		Total: len(todos),
	}

	if pages.Page != 1 {
		t.Errorf("Unexpected page number: got %v, want 1", pages.Page)
	}
	if pages.Size != 10 {
		t.Errorf("Unexpected page size: got %v, want 10", pages.Size)
	}
	if pages.Total != 1 {
		t.Errorf("Unexpected total todos: got %v, want 1", pages.Total)
	}
	if len(pages.Todos) != 1 {
		t.Errorf("Unexpected number of todos in the page: got %v, want 1", len(pages.Todos))
	}
}
