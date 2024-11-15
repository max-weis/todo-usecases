package remove_test

import (
	"context"
	"errors"
	"github.com/max-weis/todo-usecases/remove"
	"testing"
)

func TestDeleteTodoUseCase(t *testing.T) {
	dbError := errors.New("database error")

	tests := []struct {
		name        string
		id          string
		deleteErr   error
		expectedErr error
	}{
		{"Empty ID", "", nil, remove.ErrNoID},
		{"Valid ID", "todo-1", nil, nil},
		{"Delete Error", "todo-1", dbError, dbError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteFunc := func(ctx context.Context, id string) error {
				if id != tt.id {
					t.Errorf("DeleteTodoUseCase called with wrong ID: got %s, want %s", id, tt.id)
				}
				return tt.deleteErr
			}
			useCase := remove.NewRemoveTodoUseCase(deleteFunc)
			err := useCase(context.Background(), tt.id)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("DeleteTodoUseCase() error = %v, wantErr %v", err, tt.expectedErr)
			}
		})
	}
}
