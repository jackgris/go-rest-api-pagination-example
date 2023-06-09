package v1

import (
	"github.com/gofiber/fiber/v2"
)

func GetTodos(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		return c.SendString("Not implemented GetTodos")
	}

	return fn
}

func CreateTodo(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		return c.SendString("Not implemented CreateTodo")
	}

	return fn
}

func UpdateTodo(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		return c.SendString("Not implemented UpdateTodo")
	}

	return fn
}

func GetTodoById(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		return c.SendString("Not implemented GetTodoById")
	}

	return fn
}

func DeleteTodo(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		return c.SendString("Not implemented DeleteTodo")
	}

	return fn
}
