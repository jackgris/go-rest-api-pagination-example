package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		tds, err := cfg.Core.Query(c.Context(), "", "", 0, 0)
		if err != nil {
			return err
		}

		message := fmt.Sprintf("All todos %s", tds)
		return c.SendString(message)
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
