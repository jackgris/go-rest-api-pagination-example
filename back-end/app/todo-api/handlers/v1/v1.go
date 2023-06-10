package v1

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

type Config struct {
	Log  *log.Logger
	Core *todo.Core
}

func Routes(app *fiber.App, cfg Config) {
	const version = "v1"

	app.Route(version, func(api fiber.Router) {
		api.Get("/todos", GetTodos(cfg))
		api.Post("/todos", CreateTodo(cfg))
		api.Put("/todos/:id", UpdateTodo(cfg))
		api.Get("/todos/:id", GetTodoById(cfg))
		api.Delete("/todos/:id", DeleteTodo(cfg))
	}, version+".")
}
