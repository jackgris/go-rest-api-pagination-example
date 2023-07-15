package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	v1 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v1"
	v2 "github.com/jackgris/go-rest-api-pagination-example/app/todo-api/handlers/v2"
	"github.com/jackgris/go-rest-api-pagination-example/business/logger"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo"
	"github.com/jackgris/go-rest-api-pagination-example/business/todo/store/tododb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DbHost     string `yaml:"DbHost"`
	DbName     string `yaml:"DbName"`
	DbPass     string `yaml:"DbPass"`
	DbUser     string `yaml:"DbUser"`
	ServerPort string `yaml:"ServerPort"`
	DbURI      string
}

var (
	dbHost     = "localhost:3306"
	dbName     = "todo_pagination"
	dbPass     = "1234"
	dbUser     = "pagination"
	serverPort = "3000"
)

func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{
		DbHost:     dbHost,
		DbName:     dbName,
		DbPass:     dbPass,
		DbUser:     dbUser,
		ServerPort: serverPort,
	}

	// update config values from env, if any
	cfg.GETENVs()
	// init db conn
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbName)
	cfg.DbURI = dsn

	return cfg, err
}

func (c *Config) GETENVs() {
	if val, found := os.LookupEnv("CONFIG_DBHOST"); found {
		c.DbHost = val
	}

	if val, found := os.LookupEnv("CONFIG_DBNAME"); found {
		c.DbName = val
	}

	if val, found := os.LookupEnv("CONFIG_DBPASS"); found {
		c.DbPass = val
	}

	if val, found := os.LookupEnv("CONFIG_DBUSER"); found {
		c.DbUser = val
	}

	if val, found := os.LookupEnv("CONFIG_SERVER_PORT"); found {
		c.ServerPort = val
	}
}

// NewApp create the application server
func NewApp(core *todo.Core, logger *logger.Logger) *fiber.App {

	cfg := v1.Config{
		Log:  logger,
		Core: core,
	}
	cfg2 := v2.Config{
		Log:  logger,
		Core: core,
	}

	app := fiber.New()
	// Initialize default config
	app.Use(cors.New())

	v1.Routes(app, cfg)
	v2.Routes(app, cfg2)

	return app
}

// NewDb create the database connection
func NewDb(log *logger.Logger) *todo.Core {

	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Can't get config: %s", err)
	}

	log.Println("Checking database")
	err = checkDb(log, *cfg)
	if err != nil {
		log.Fatalf("Failed to create database: %s", err)
	}
	log.Println("Database it's fine")

	log.Println("Connecting to database")
	db, err := gorm.Open(mysql.Open(cfg.DbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}
	log.Println("Connection successfully")

	log.Println("Running Migrations")
	err = db.AutoMigrate(&tododb.DbTodo{})
	if err != nil {
		log.Fatalf("Failed to run migration database: %s", err)
	}
	log.Println("Migration successfully")

	store := tododb.NewStore(log, db)
	core := todo.NewCore(log, store)

	return core
}

// checkDb is used because this application is not for production, so can happen that the database don't exist
func checkDb(log *logger.Logger, cfg Config) error {
	db, err := gorm.Open(mysql.Open(cfg.DbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database while checking state: %s", err)
	}
	defer func() {
		instance, _ := db.DB()
		_ = instance.Close()
	}()

	count := 0
	db.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", cfg.DbName).Scan(&count)
	if count == 0 {
		sql := fmt.Sprintf("CREATE DATABASE %s", cfg.DbName)
		result := db.Exec(sql)
		return result.Error
	}

	return nil
}
