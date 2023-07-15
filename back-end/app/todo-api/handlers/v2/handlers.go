package v2

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

func GetTodos(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {

		var offset int
		var limit int
		var err error
		// Read the values from url
		_offset := c.Query("offset", "1")
		_limit := c.Query("limit", "10")
		// Validate the data
		offset, err = strconv.Atoi(_offset)
		if err != nil {
			offset = 1
		}
		limit, err = strconv.Atoi(_limit)
		if err != nil {
			limit = 10
		}
		if limit > 100 {
			limit = 100
		}

		tds, err := cfg.Core.Query(c.Context(), "", "", offset, limit)
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

		num, err := cfg.Core.Count(c.Context(), "")
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(Response{
				Success: false,
				Message: "Can't proccess number of Todos",
			})
		}
		// Calculate the number of pages for our fron-end
		pages := 1
		if int(num) > limit {
			pages = (int(num) / limit) + 1
		}
		c.Status(fiber.StatusOK)
		return c.JSON(Response{
			Success: true,
			Data:    tdsJson,
			Pages:   pages,
		})
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
			Title:       td.Title,
			Description: td.Description,
		}
		fmt.Println("NEW TODO: ", newTd)
		dbTd, err := cfg.Core.Create(c.Context(), newTd)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(Response{
				Success: false,
				Message: "Can't create Todo in database",
			})
		}
		c.Status(fiber.StatusCreated)
		todos := []Todo{}
		todos = append(todos, todoToTodoJson(dbTd))
		return c.JSON(Response{
			Success: true,
			Data:    todos,
		})
	}

	return fn
}
