package main

func main() {
	const port = ":3000"
	// Configure our logger
	logger := NewLogger()
	// Create a new database connection
	db := NewDb(logger)
	// Create our server and run
	app := NewApp(db, logger)
	logger.Fatal(app.Listen(port))
}
