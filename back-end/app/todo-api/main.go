package main

import "github.com/jackgris/go-rest-api-pagination-example/business/logger"

func main() {

	const port = ":3000"
	// Configure our logger
	logger := logger.New()
	// Create a new database connection
	db := NewDb(logger)
	// Create our server and run
	app := NewApp(db, logger)
	logger.Fatal(app.Listen(port))
}
