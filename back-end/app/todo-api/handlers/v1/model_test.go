package v1_test

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

// MockStore have the purpose of simulate our database data
type MockStore struct {
	data map[string]todo.Todo
}

// NewStore constructs the api for data access.
func NewStore(data map[string]todo.Todo) *MockStore {
	return &MockStore{
		data: data,
	}
}

func (m MockStore) Create(ctx context.Context, todo todo.Todo) error {
	m.data[todo.ID.String()] = todo
	return nil
}

func (m MockStore) Update(ctx context.Context, todo todo.Todo) error {
	m.data[todo.ID.String()] = todo
	return nil
}

func (m MockStore) Delete(ctx context.Context, todo todo.Todo) error {
	delete(m.data, todo.ID.String())
	return nil
}

func (m MockStore) Query(ctx context.Context, filter string, orderBy string, pageNumber int, rowsPerPage int) ([]todo.Todo, error) {
	count := 0
	amount := 0
	amountTodos := 1
	if pageNumber > 1 {
		amountTodos = (pageNumber - 1) * rowsPerPage
	}

	todos := []todo.Todo{}
	for _, t := range m.data {
		if amount >= amountTodos {
			if count >= rowsPerPage {
				break
			}

			todos = append(todos, t)
			count++
		}
		amount++
	}

	return todos, nil
}

func (m MockStore) Count(ctx context.Context, filter string) (int64, error) {
	count := int64(len(m.data))
	return count, nil
}

func (m MockStore) QueryByID(ctx context.Context, todoID uuid.UUID) (todo.Todo, error) {
	td, ok := m.data[todoID.String()]
	if !ok {
		return td, errors.New("Not found")
	}
	return td, nil
}
