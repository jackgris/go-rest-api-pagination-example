package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v1"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo/store/tododb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewApp create the application server
func NewApp(core *todo.Core, logger *logger.Logger) *fiber.App {

	cfg := v1.Config{
		Log:  logger,
		Core: core,
	}
	app := fiber.New()
	v1.Routes(app, cfg)

	return app
}

// NewDb create the database connection
func NewDb(log *logger.Logger) *todo.Core {
	dbname := "todo_pagination"

	log.Println("Checking database")
	err := checkDb(log, dbname)
	if err != nil {
		log.Fatalf("Failed to create database: %s", err)
	}
	log.Println("Database it's fine")

	log.Println("Connecting to database")
	dsn := "pagination:1234@tcp(172.17.0.2:3306)/" + dbname
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}
	log.Println("Connection successfully")

	log.Println("Running Migrations")
	err = db.AutoMigrate(todo.Todo{})
	if err != nil {
		log.Fatalf("Failed to run migration database: %s", err)
	}
	log.Println("Migration successfully")

	store := tododb.NewStore(log, db)
	core := todo.NewCore(log, store)

	return core
}

// checkDb is used because this application is not for production, so can happen that the database don't exist
func checkDb(log *logger.Logger, dbname string) error {
	dsn := "pagination:1234@tcp(172.17.0.2:3306)/"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database while checking state: %s", err)
	}
	defer func() {
		instance, _ := db.DB()
		_ = instance.Close()
	}()

	count := 0
	db.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbname).Scan(&count)
	if count == 0 {
		sql := fmt.Sprintf("CREATE DATABASE %s", dbname)
		result := db.Exec(sql)
		return result.Error
	}

	return nil
}
