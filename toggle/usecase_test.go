package toggle_test

import (
	"context"
	"errors"
	"github.com/max-weis/todo-usecases/toggle"
	"testing"
)

func TestToggleTodosUseCase(t *testing.T) {
	dbError := errors.New("database error")
	tests := []struct {
		name        string
		id          string
		toggleErr   error
		expectedErr error
	}{
		{"Empty ID", "", nil, toggle.ErrNoID},
		{"Valid ID", "todo-1", nil, nil},
		{"Toggle Error", "todo-1", dbError, dbError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toggleFunc := func(ctx context.Context, id string) error {
				if id != tt.id {
					t.Errorf("ToggleTodosUseCase called with wrong ID: got %s, want %s", id, tt.id)
				}
				return tt.toggleErr
			}
			useCase := toggle.NewToggleTodosUseCase(toggleFunc)
			err := useCase(context.Background(), tt.id)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("ToggleTodosUseCase() error = %v, wantErr %v", err, tt.expectedErr)
			}
		})
	}
}
