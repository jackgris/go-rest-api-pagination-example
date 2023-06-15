package tododb

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
	"gorm.io/gorm"
)

var ErrNotID = errors.New("todo need ID")
var ErrNotExist = errors.New("todo ID don't found")

// Store manages the set of APIs for todo database access.
type Store struct {
	log *logger.Logger
	db  *gorm.DB
}

// NewStore constructs the api for data access.
func NewStore(log *logger.Logger, db *gorm.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// Create adds a Todo to the database. It returns the created Todo with
// fields like ID and DateCreated populated.
func (s *Store) Create(ctx context.Context, td todo.Todo) error {
	dbTd := toDBTodo(td)
	result := s.db.Create(&dbTd) // pass pointer of data to Create

	return result.Error
}

func (s *Store) Update(ctx context.Context, td todo.Todo) error {
	if td.ID == uuid.Nil {
		return ErrNotID
	}
	dbTd := toDBTodo(td)
	result := s.db.First(&dbTd)
	// Check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ErrNotExist
	}

	result = s.db.Save(&dbTd)

	return result.Error
}

func (s *Store) Delete(ctx context.Context, td todo.Todo) error {
	dbTd := toDBTodo(td)
	result := s.db.Delete(&DbTodo{}, dbTd.ID)

	return result.Error
}

func (s *Store) Query(ctx context.Context, filter string, orderBy string, pageNumber int, rowsPerPage int) ([]todo.Todo, error) {

	dbTd := []DbTodo{}
	// Set minimun number of page
	if pageNumber <= 0 {
		pageNumber = 1
	}

	// Set minimun and maximum number of row per page
	switch {
	case rowsPerPage > 100:
		rowsPerPage = 100
	case rowsPerPage <= 0:
		rowsPerPage = 10
	}

	// Start showing result from this number of Todos
	offset := (pageNumber - 1) * rowsPerPage

	result := s.db.Offset(offset).Limit(rowsPerPage).Find(&dbTd)
	if result.Error != nil {
		return []todo.Todo{}, result.Error
	}

	return toTodoSlice(dbTd), nil
}

func (s *Store) Count(ctx context.Context, filter string) (int64, error) {
	var count int64

	result := s.db.Model(&DbTodo{}).Count(&count)
	if result.Error != nil {
		return 0, nil
	}

	return count, nil
}

func (s *Store) QueryByID(ctx context.Context, todoID uuid.UUID) (todo.Todo, error) {
	var td = DbTodo{ID: todoID}
	result := s.db.First(&td)

	if result.Error != nil {
		return todo.Todo{}, nil
	}
	return toCoreTodo(td), nil
}
