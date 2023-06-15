package v1_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v1"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
)

var todos []v1.Todo

func TestMain(m *testing.M) {
	var err error
	todos, err = getData()
	if err != nil {
		fmt.Println(err)
	}

	m.Run()
}

func getData() ([]v1.Todo, error) {

	tds := []v1.Todo{}

	// Open our jsonFile
	jsonFile, err := os.Open("../../../../data/todos.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return tds, err
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return tds, err
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &tds)
	if err != nil {
		return tds, err
	}

	return tds, nil
}

func todoJsonToTodo(td v1.Todo) todo.Todo {
	return todo.Todo{
		ID:          td.ID,
		Name:        td.Name,
		Description: td.Description,
		DateCreated: td.DateCreated,
		DateUpdated: td.DateUpdated,
	}
}

func newDb(log *logger.Logger) *todo.Core {
	data := map[string]todo.Todo{}

	for _, t := range todos {
		newT := todoJsonToTodo(t)
		data[newT.ID.String()] = newT
	}

	store := NewStore(data)
	core := todo.NewCore(log, store)

	return core
}

// NewApp create the application server
func newApp(core *todo.Core, logger *logger.Logger) *fiber.App {

	cfg := v1.Config{
		Log:  logger,
		Core: core,
	}
	app := fiber.New()
	v1.Routes(app, cfg)

	return app
}
