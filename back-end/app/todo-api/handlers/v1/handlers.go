package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

var numTodos int = 10

func GetTodos(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		// Get the page number, if the user don't pass any number the default will be the first page
		page := c.Query("page", "1")
		pageNumber, err := strconv.Atoi(page)
		if err != nil {
			pageNumber = 1
		}
		// We search all Todos related to the number page and with a maximum per page of 10 results
		tds, err := cfg.Core.Query(c.Context(), "", "", pageNumber, numTodos)
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
		if int(num) > numTodos {
			pages = (int(num) / numTodos) + 1
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

func UpdateTodo(cfg Config) fiber.Handler {

	fn := func(c *fiber.Ctx) error {
		tdJson, err := postToTodo(c.Body())
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(Response{
				Success: false,
				Message: "Can't proccess Todo",
			})
		}

		if tdJson.ID == uuid.Nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(Response{
				Success: false,
				Message: "It's not a valid Todo",
			})
		}
		uTd := todo.UpdateTodo{}
		if tdJson.Title != "" {
			uTd.Title = &tdJson.Title
		}
		if tdJson.Description != "" {
			uTd.Description = &tdJson.Description
		}

		uTd.Completed = &tdJson.Completed

		td := todoJsonToTodo(tdJson)
		td, err = cfg.Core.Update(c.Context(), td, uTd)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(Response{
				Success: false,
				Message: "Can't update Todo",
			})
		}
		c.Status(fiber.StatusOK)
		todos := []Todo{}
		todos = append(todos, todoToTodoJson(td))
		return c.JSON(Response{
			Success: true,
			Data:    todos,
		})
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
		if err != nil || td.ID == uuid.Nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(Response{
				Success: false,
				Message: "ID not found",
			})
		}

		todos := []Todo{}
		todos = append(todos, todoToTodoJson(td))
		c.Status(fiber.StatusOK)
		return c.JSON(Response{
			Success: true,
			Message: "",
			Data:    todos,
		})
	}

	return fn
}

func DeleteTodo(cfg Config) fiber.Handler {

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
		td := todo.Todo{ID: uuId}
		err = cfg.Core.Delete(c.Context(), td)
		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(Response{
				Success: false,
				Message: "ID not found",
			})
		}

		c.Status(fiber.StatusOK)
		return c.JSON(Response{
			Success: true,
			Message: "Todo deleted",
		})
	}

	return fn
}
