package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/photkosee/todo-go/handlers"
)

func TodoRoutes(app *fiber.App) {
	todo := app.Group("/api/v1/todos")
	todo.Get("/", handlers.GetTodos)
	todo.Post("/", handlers.CreateTodo)
	todo.Patch("/:id", handlers.UpdateTodo)
	todo.Patch("/:id/complete", handlers.CompleteTodo)
	todo.Delete("/:id", handlers.DeleteTodo)
}
