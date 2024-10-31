package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/photkosee/todo-go/config"
	"github.com/photkosee/todo-go/database"
	"github.com/photkosee/todo-go/routes"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		return
	}

	// start database
	err = database.StartMongoDB()
	if err != nil {
		return
	}

	// defer closing database
	defer database.CloseMongoDB()

	app := fiber.New()

	// setup routes
	routes.TodoRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen("0.0.0.0:" + port))

}
