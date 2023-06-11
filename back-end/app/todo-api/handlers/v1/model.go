package v1

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

// Todo represents an individual todo.
type Todo struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// Response represents the server information send
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Todo   `json:"data"`
}

func postToTodo(data []byte) (todo.Todo, error) {

	td := todo.Todo{}
	err := json.Unmarshal(data, &td)
	if err != nil {
		return todo.Todo{}, err
	}

	return td, nil
}

func todoToTodoJson(td todo.Todo) Todo {
	return Todo{
		ID:          td.ID,
		Name:        td.Name,
		Description: td.Description,
		DateCreated: td.DateCreated,
		DateUpdated: td.DateUpdated,
	}
}

func isNotValidTodo(td todo.Todo) bool {
	return (td.Name == "" || td.Description == "")
}
