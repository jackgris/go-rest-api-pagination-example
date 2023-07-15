package v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

type Config struct {
	Log  *logger.Logger
	Core *todo.Core
}

func Routes(app *fiber.App, cfg Config) {
	const version = "v2"

	app.Route(version, func(api fiber.Router) {
		api.Get("/todos", GetTodos(cfg))
		api.Post("/todos", CreateTodo(cfg))
		// api.Put("/todos", UpdateTodo(cfg))
		// api.Get("/todos/:id", GetTodoById(cfg))
		// api.Delete("/todos/:id", DeleteTodo(cfg))
	}, version+".")
}
