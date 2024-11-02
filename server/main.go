package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	// setup routes
	routes.TodoRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen("0.0.0.0:" + port))

}
