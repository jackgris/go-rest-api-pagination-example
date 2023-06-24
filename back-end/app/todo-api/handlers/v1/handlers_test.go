package v1_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	v1 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v1"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
)

func TestGetTodos(t *testing.T) {
	logger := logger.New()
	db := newDb(logger)
	app := newApp(db, logger)
	params := []Params{
		{url: "/v1/todos?page=1", code: fiber.StatusOK, method: fiber.MethodGet, amount: 10},
		{url: "/v1/todos?page=3", code: fiber.StatusOK, method: fiber.MethodGet, amount: 10},
		{url: "/v1/todos?page=4", code: fiber.StatusOK, method: fiber.MethodGet, amount: 5},
		{url: "/v1/todos?page=0", code: fiber.StatusOK, method: fiber.MethodGet, amount: 10},
		{url: "/v1/todos", code: fiber.StatusOK, method: fiber.MethodGet, amount: 10},
		{url: "/v1/todos?", code: fiber.StatusOK, method: fiber.MethodGet, amount: 10},
		{url: "/v1/todos?page=", code: fiber.StatusOK, method: fiber.MethodGet, amount: 10},
		{url: "/todos?page=2", code: fiber.StatusNotFound, method: fiber.MethodGet, result: nil, amount: 0},
		{url: "/todos?page=2", code: fiber.StatusNotFound, method: fiber.MethodPost, result: nil, amount: 0},
	}

	for _, p := range params {
		runningApp(app, t, p, logger)
	}
}

func TestCreateTodos(t *testing.T) {
	logger := logger.New()
	db := newDb(logger)
	app := newApp(db, logger)
	params := []Params{
		{url: "/v1/todos", code: fiber.StatusBadRequest, method: fiber.MethodPost, result: nil, amount: 0},
		{url: "/v1/todos", code: fiber.StatusCreated, method: fiber.MethodPost,
			send:   &v1.Todo{Title: "Testing", Description: "Some data"},
			result: &v1.Todo{Title: "Testing", Description: "Some data"}},
	}

	for _, p := range params {
		runningApp(app, t, p, logger)
	}
}

func TestUpdateTodo(t *testing.T) {
	logger := logger.New()
	db := newDb(logger)
	app := newApp(db, logger)
	// You can find this on the data folder inside the json file
	id, err := uuid.Parse("f8f57b2a-4c3d-48fc-bae9-7f36568de9dc")
	if err != nil {
		t.Fatal(err)
	}

	fakeId := uuid.New()

	params := []Params{
		{url: "/v1/todos", code: fiber.StatusOK, method: fiber.MethodPut,
			send:   &v1.Todo{ID: id, Title: "New name", Description: "New description"},
			result: &v1.Todo{ID: id, Title: "New name", Description: "New description"}},
		{url: "/v1/todos", code: fiber.StatusInternalServerError, method: fiber.MethodPut,
			send:   &v1.Todo{ID: fakeId, Title: "New name", Description: "New description"},
			result: &v1.Todo{ID: fakeId, Title: "", Description: ""}},
	}

	for _, p := range params {
		runningApp(app, t, p, logger)
	}
}

func TestGetByIdTodo(t *testing.T) {
	logger := logger.New()
	db := newDb(logger)
	app := newApp(db, logger)
	// You can find this on the data folder inside the json file
	id, err := uuid.Parse("f8f57b2a-4c3d-48fc-bae9-7f36568de9dc")
	if err != nil {
		t.Fatal(err)
	}

	fakeId := uuid.New()

	params := []Params{
		{url: "/v1/todos/" + id.String(), code: fiber.StatusOK, method: fiber.MethodGet, byID: true,
			result: &v1.Todo{ID: id, Title: "Another 2aaa22", Description: "Chau"}},
		{url: "/v1/todos/" + fakeId.String(), code: fiber.StatusNotFound, method: fiber.MethodGet,
			byID: true, result: &v1.Todo{ID: fakeId, Title: "", Description: ""}},
	}

	for _, p := range params {
		runningApp(app, t, p, logger)
	}
}

func TestDeleteTodo(t *testing.T) {
	logger := logger.New()
	db := newDb(logger)
	app := newApp(db, logger)
	// You can find this on the data folder inside the json file
	id, err := uuid.Parse("f8f57b2a-4c3d-48fc-bae9-7f36568de9dc")
	if err != nil {
		t.Fatal(err)
	}

	fakeId := uuid.New()

	params := []Params{
		{url: "/v1/todos/" + fakeId.String(), code: fiber.StatusNotFound, method: fiber.MethodDelete},
		{url: "/v1/todos/" + id.String(), code: fiber.StatusOK, method: fiber.MethodDelete},
		{url: "/v1/todos/" + id.String(), code: fiber.StatusNotFound, method: fiber.MethodDelete},
	}

	for _, p := range params {
		runningApp(app, t, p, logger)
	}
}

type Params struct {
	url    string
	code   int
	method string
	result *v1.Todo
	amount int
	send   *v1.Todo
	byID   bool
}

func runningApp(app *fiber.App, t *testing.T, p Params, l *logger.Logger) {

	var buf bytes.Buffer
	if p.send != nil {
		err := json.NewEncoder(&buf).Encode(p.send)
		if err != nil {
			l.Printf("Can't create body data: %v", p.send)
			t.Fail()
		}
	}

	resp, err := app.Test(httptest.NewRequest(p.method, p.url, &buf))
	if err != nil {
		l.Println("Error while run app.Test: ", err)
		t.Fail()
	}
	if resp.StatusCode != p.code {
		l.Printf("Status code should be %d but receive %d\n", p.code, resp.StatusCode)
		t.Fail()
	}

	if p.result != nil {
		byteValue, err := io.ReadAll(resp.Body)
		if err != nil {
			l.Println("Can't read body response: ", err)
			t.Fail()
		}

		if p.method == fiber.MethodGet {
			if !p.byID {
				checkGetAll(byteValue, l, p, t)
			}
		} else {
			check(byteValue, l, p, t)
		}
	}
}

func checkGetAll(byteValue []byte, l *logger.Logger, p Params, t *testing.T) {

	tds := []v1.Todo{}
	err := json.Unmarshal(byteValue, &tds)
	if err != nil {
		l.Println("Can't unmarshal response: ", err)
		t.Fail()
	}

	if len(tds) != p.amount {
		l.Printf("The amount of todos should be %d but receive %d: ", p.amount, len(tds))
		t.Fail()
	}
}

func check(byteValue []byte, l *logger.Logger, p Params, t *testing.T) {

	r := v1.Response{}
	err := json.Unmarshal(byteValue, &r)
	if err != nil {
		l.Println("Can't unmarshal response: ", err)
		t.Fail()
	}

	if len(r.Data) > 0 {
		if r.Data[0].Title != p.result.Title {
			l.Printf("We receive name %s, but want %s\n", r.Data[0].Title, p.result.Title)
			t.Fail()
		}

		if r.Data[0].Description != p.result.Description {
			l.Printf("We receive description %s, but want %s\n", r.Data[0].Description, p.result.Description)
			t.Fail()
		}
	}
}
