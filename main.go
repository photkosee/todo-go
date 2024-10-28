package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
	})

	// get all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// create todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := new(Todo)

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)
	})

  // Update a todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

    payload := struct {
      Body string `json:"body"`
    }{}

    if err := c.BodyParser(&payload); err != nil {
      return err
    }

    if payload.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Body = payload.Body
				return c.Status(200).JSON(todos)
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// complete a todo
	app.Patch("/api/todos/:id/complete", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos)
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

  // delete a todo
  app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
    id := c.Params("id")

    for i, todo := range todos {
      if fmt.Sprint(todo.ID) == id {
        todos = append(todos[:i], todos[i+1:]...)
        return c.Status(200).JSON(todos)
      }
    }

    return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
  })

	log.Fatal(app.Listen(":3005"))
}
