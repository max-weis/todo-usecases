package todo_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/max-weis/todo-usecases/create"
)

var now = time.Now()

// TestTodo_Validate tests the Validate method of the Todo struct.
func TestTodo_Validate(t *testing.T) {
	tests := []struct {
		name    string
		todo    todo.Todo
		wantErr error
	}{
		{"Empty Title", todo.Todo{Title: "", DueDate: now.Add(24 * time.Hour)}, todo.ErrEmptyTitle},
		{"Valid Todo", todo.Todo{Title: "Buy Milk", DueDate: now.Add(24 * time.Hour)}, nil},
		{"Past Due Date", todo.Todo{Title: "Expired Task", DueDate: now.Add(-24 * time.Hour)}, todo.ErrInvalidDueDate},
		{"Whitespace Title", todo.Todo{Title: "   ", DueDate: now.Add(24 * time.Hour)}, todo.ErrEmptyTitle},
		{"No Due Date", todo.Todo{Title: "No Due Date"}, todo.ErrInvalidDueDate},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.todo.Validate(); !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestCreateTodoUseCase tests the returned CreateTodoUseCase function.
func TestCreateTodoUseCase(t *testing.T) {
	var dbError = errors.New("database error")
	tests := []struct {
		name          string
		title         string
		dueDate       time.Time
		saveTodoErr   error
		expectedErr   error
		expectedTitle string
	}{
		{"Valid Creation", "Do Homework", now.Add(24 * time.Hour), nil, nil, "Do Homework"},
		{"Empty Title", "", now.Add(24 * time.Hour), nil, todo.ErrEmptyTitle, ""},
		{"Past Due Date", "Late Task", now.Add(-24 * time.Hour), nil, todo.ErrInvalidDueDate, ""},
		{"Save Error", "Save Error", now.Add(24 * time.Hour), dbError, dbError, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveTodo := func(ctx context.Context, todo todo.Todo) error {
				return tt.saveTodoErr
			}
			useCase := todo.NewCreateTodoUseCase(saveTodo)
			_, err := useCase(context.Background(), tt.title, tt.dueDate)

			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("CreateTodoUseCase() error = %v, wantErr %v", err, tt.expectedErr)
				}
			} else if err != nil {
				t.Errorf("CreateTodoUseCase() unexpected error: %v", err)
			}
		})
	}
}
