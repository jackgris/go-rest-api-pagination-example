package main

func main() {
	const port = ":3000"
	// Create a new database connection
	db := NewDb()
	// Configure our logger
	logger := NewLogger()
	// Create our server and run
	app := NewApp(db, logger)
	logger.Fatal(app.Listen(port))
}
