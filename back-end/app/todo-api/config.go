package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v1"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo/store/tododb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewApp create the application server
func NewApp(core *todo.Core, logger *log.Logger) *fiber.App {

	cfg := v1.Config{
		Log:  logger,
		Core: core,
	}
	app := fiber.New()
	v1.Routes(app, cfg)

	return app
}

// NewDb create the database connection
func NewDb(log *log.Logger) *todo.Core {

	dsn := "pagination:1234@tcp(172.17.0.2:3306)/todo_pagination"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	store := tododb.NewStore(log, db)
	core := todo.NewCore(log, store)

	return core
}

// NewLogger configure our logs
func NewLogger() *log.Logger {

	logger := log.Logger{}
	return &logger
}
