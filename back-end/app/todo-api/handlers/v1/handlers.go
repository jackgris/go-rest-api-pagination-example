package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

func GetTodos(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		tds, err := cfg.Core.Query(c.Context(), "", "", 0, 0)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(Response{
				Success: false,
				Message: "Can't proccess Todos",
			})
		}
		tdsJson := []Todo{}
		for _, t := range tds {
			tdsJson = append(tdsJson, todoToTodoJson(t))
		}
		c.Status(fiber.StatusOK)
		return c.JSON(tdsJson)
	}

	return fn
}

func CreateTodo(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {

		td, err := postToTodo(c.Body())
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(Response{
				Success: false,
				Message: "Can't proccess Todo",
			})
		}

		if isNotValidTodo(td) {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(Response{
				Success: false,
				Message: "Todo should contain name and description",
			})
		}

		newTd := todo.NewTodo{
			Name:        td.Name,
			Description: td.Description,
		}

		dbTd, err := cfg.Core.Create(c.Context(), newTd)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(Response{
				Success: false,
				Message: "Can't create Todo in database",
			})
		}
		c.Status(fiber.StatusCreated)
		return c.JSON(Response{
			Success: true,
			Data:    todoToTodoJson(dbTd),
		})
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
		id := c.Params("id", "")
		if id == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(Response{
				Success: false,
				Message: "You need to pass an ID params",
			})
		}
		uuId, err := uuid.Parse(id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(Response{
				Success: false,
				Message: "Invalid ID params",
			})
		}
		td, err := cfg.Core.QueryByID(c.Context(), uuId)
		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(Response{
				Success: false,
				Message: "ID not found",
			})
		}
		tdJson := todoToTodoJson(td)

		c.Status(fiber.StatusOK)
		return c.JSON(Response{
			Success: true,
			Message: "",
			Data:    tdJson,
		})
	}

	return fn
}

func DeleteTodo(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		return c.SendString("Not implemented DeleteTodo")
	}

	return fn
}
