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
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
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
