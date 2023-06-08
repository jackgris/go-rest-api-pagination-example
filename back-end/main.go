package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	dsn := "pagination:1234@tcp(172.17.0.2:3306)/todo_pagination"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		panic("Can't migrate schema")
	}

	app := fiber.New()

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.SendString("Future todo list with pagination")
	})

	log.Fatal(app.Listen(":3000"))
}
