package v2

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

// Todo represents an individual todo.
type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}

// Response represents the server information send
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []Todo `json:"data"`
	Pages   int    `json:"pages"`
}

func postToTodo(data []byte) (Todo, error) {

	td := Todo{}
	err := json.Unmarshal(data, &td)
	if err != nil {
		return Todo{}, err
	}

	return td, nil
}

func todoToTodoJson(td todo.Todo) Todo {
	return Todo{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Completed:   td.Completed,
	}
}

// func todoJsonToTodo(td Todo) todo.Todo {
// 	return todo.Todo{
// 		ID:          td.ID,
// 		Title:       td.Title,
// 		Completed:   td.Completed,
// 		Description: td.Description,
// 	}
// }

func isNotValidTodo(td Todo) bool {
	return (td.Title == "" || td.Description == "")
}
