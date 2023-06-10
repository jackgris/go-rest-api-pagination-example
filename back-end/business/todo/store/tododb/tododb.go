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
	result := s.db.Save(&dbTd)

	return result.Error
}

func (s *Store) Delete(ctx context.Context, td todo.Todo) error {
	dbTd := toDBTodo(td)
	result := s.db.Delete(&dbTodo{}, dbTd.ID)

	return result.Error
}

func (s *Store) Query(ctx context.Context, filter string, orderBy string, pageNumber int, rowsPerPage int) ([]todo.Todo, error) {

	dbTd := []dbTodo{}
	// Get all records
	result := s.db.Find(&dbTd)

	if result.Error != nil {
		return []todo.Todo{}, nil
	}
	return toTodoSlice(dbTd), nil
}

func (s *Store) Count(ctx context.Context, filter string) (int64, error) {
	var count int64

	result := s.db.Model(&dbTodo{}).Count(&count)
	if result.Error != nil {
		return 0, nil
	}

	return count, nil
}

func (s *Store) QueryByID(ctx context.Context, todoID uuid.UUID) (todo.Todo, error) {
	var td = dbTodo{ID: todoID}
	result := s.db.First(&td)

	if result.Error != nil {
		return todo.Todo{}, nil
	}
	return toCoreTodo(td), nil
}