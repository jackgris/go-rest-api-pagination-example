package tododb

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

// Todo represents an individual todo in the database.
type DbTodo struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	DateCreated time.Time
	DateUpdated time.Time
}

// =============================================================================

func toDBTodo(td todo.Todo) DbTodo {
	todoDB := DbTodo{
		ID:          td.ID,
		Name:        td.Name,
		Description: td.Description,
		DateCreated: td.DateCreated.UTC(),
		DateUpdated: td.DateUpdated.UTC(),
	}

	return todoDB
}

func toCoreTodo(dbTd DbTodo) todo.Todo {
	td := todo.Todo{
		ID:          dbTd.ID,
		Name:        dbTd.Name,
		Description: dbTd.Description,
		DateCreated: dbTd.DateCreated.In(time.Local),
		DateUpdated: dbTd.DateUpdated.In(time.Local),
	}

	return td
}

func toTodoSlice(dbTodos []DbTodo) []todo.Todo {
	todos := make([]todo.Todo, len(dbTodos))
	for i, dbTd := range dbTodos {
		todos[i] = toCoreTodo(dbTd)
	}
	return todos
}
