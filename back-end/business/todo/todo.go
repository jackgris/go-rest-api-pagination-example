package todo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
)

// Set of error variables for CRUD operations.
var ErrNotFound = errors.New("todo not found")

// =============================================================================

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, todo Todo) error
	Update(ctx context.Context, todo Todo) error
	Delete(ctx context.Context, todo Todo) error
	Query(ctx context.Context, filter string, orderBy string, pageNumber int, rowsPerPage int) ([]Todo, error)
	Count(ctx context.Context, filter string) (int64, error)
	QueryByID(ctx context.Context, todoID uuid.UUID) (Todo, error)
}

// Core manages the set of APIs for todo access.
type Core struct {
	log    *logger.Logger
	storer Storer
}

// NewCore constructs a core for todo api access.
func NewCore(log *logger.Logger, storer Storer) *Core {
	core := Core{
		log:    log,
		storer: storer,
	}

	return &core
}

// Create adds a Todo to the database. It returns the created Todo with
// fields like ID and DateCreated populated.
func (c *Core) Create(ctx context.Context, todo NewTodo) (Todo, error) {

	now := time.Now()

	td := Todo{
		ID:          uuid.New(),
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   false,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, td); err != nil {
		return Todo{}, fmt.Errorf("create: %w", err)
	}

	return td, nil
}

// Update modifies data about a Todo. It will error if the specified ID is
// invalid or does not reference an existing Product.
func (c *Core) Update(ctx context.Context, td Todo, up UpdateTodo) (Todo, error) {
	if up.Title != nil {
		td.Title = *up.Title
	}
	if up.Description != nil {
		td.Description = *up.Description
	}
	if up.Completed != nil {
		td.Completed = *up.Completed
	}
	td.DateUpdated = time.Now()

	if err := c.storer.Update(ctx, td); err != nil {
		return Todo{}, fmt.Errorf("update: %w", err)
	}

	return td, nil
}

// Delete removes the todo identified by a given ID.
func (c *Core) Delete(ctx context.Context, td Todo) error {
	if err := c.storer.Delete(ctx, td); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query gets all Todos from the database.
func (c *Core) Query(ctx context.Context, filter string, orderBy string, pageNumber int, rowsPerPage int) ([]Todo, error) {
	todos, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return todos, nil
}

// Count returns the total number of Todos in the store.
func (c *Core) Count(ctx context.Context, filter string) (int64, error) {
	return c.storer.Count(ctx, filter)
}

// QueryByID finds the todo identified by a given ID.
func (c *Core) QueryByID(ctx context.Context, productID uuid.UUID) (Todo, error) {
	td, err := c.storer.QueryByID(ctx, productID)
	if err != nil {
		return Todo{}, fmt.Errorf("query: productID[%s]: %w", productID, err)
	}

	return td, nil
}
