package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewApp create the application server
func NewApp(db *gorm.DB, logger *log.Logger) *fiber.App {

	cfg := v1.Config{
		Log: logger,
		Db:  db,
	}
	app := fiber.New()
	v1.Routes(app, cfg)

	return app
}

// NewDb create the database connection
func NewDb() *gorm.DB {

	dsn := "pagination:1234@tcp(172.17.0.2:3306)/todo_pagination"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// NewLogger configure our logs
func NewLogger() *log.Logger {

	logger := log.Logger{}
	return &logger
}
