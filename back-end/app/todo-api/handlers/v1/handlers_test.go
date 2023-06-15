package v1_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v1"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
)

func TestGetTodos(t *testing.T) {
	logger := logger.New()
	db := newDb(logger)
	app := newApp(db, logger)
	p := Params{
		code:       fiber.StatusOK,
		method:     fiber.MethodGet,
		result:     v1.Todo{},
		amount:     5,
		pageNumber: 4,
	}
	runningApp(app, t, p)
}

type Params struct {
	code       int
	method     string
	result     v1.Todo
	amount     int
	pageNumber int
}

func runningApp(app *fiber.App, t *testing.T, p Params) {
	page := strconv.Itoa(p.pageNumber)

	resp, err := app.Test(httptest.NewRequest(p.method, "/v1/todos?page="+page, nil))
	if err != nil {
		t.Fail()
	}
	if resp.StatusCode != p.code {
		t.Fail()
	}

	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	tds := []v1.Todo{}

	err = json.Unmarshal(byteValue, &tds)
	if err != nil {
		t.Fail()
	}

	if len(tds) != p.amount {
		t.Fail()
	}
}
