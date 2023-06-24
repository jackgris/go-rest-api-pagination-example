package tododb

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

// Todo represents an individual todo in the database.
type DbTodo struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Title       string
	Description string
	Completed   bool
	DateCreated time.Time
	DateUpdated time.Time
}

func (DbTodo) TableName() string {
	return "todo"
}

// =============================================================================

func toDBTodo(td todo.Todo) DbTodo {
	todoDB := DbTodo{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Completed:   td.Completed,
		DateCreated: td.DateCreated.UTC(),
		DateUpdated: td.DateUpdated.UTC(),
	}

	return todoDB
}

func toCoreTodo(dbTd DbTodo) todo.Todo {
	td := todo.Todo{
		ID:          dbTd.ID,
		Title:       dbTd.Title,
		Description: dbTd.Description,
		Completed:   dbTd.Completed,
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
