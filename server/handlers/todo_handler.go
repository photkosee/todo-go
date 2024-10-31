package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/photkosee/todo-go/database"
	"github.com/photkosee/todo-go/models"
)

func GetTodos(c *fiber.Ctx) error {
	todoCollection := database.GetCollection("todos")

	// return all todos
	filter := bson.M{}

	// find all todos
	cursor, err := todoCollection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// convert cursor to slice
	todos := make([]models.Todo, 0)
	if err = cursor.All(c.Context(), &todos); err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return todos
	return c.Status(200).JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todoCollection := database.GetCollection("todos")
	todo := new(models.Todo)
	// {id:0,completed:false,body:""}

	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	todo.ID = primitive.NewObjectID()

	_, err := todoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}


	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	todoCollection := database.GetCollection("todos")
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}

	var updateData struct {
		Body     string `json:"body"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	update := bson.M{"$set": bson.M{"Body": updateData.Body}}
	_, err = todoCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.SendStatus(200)
}

func CompleteTodo(c *fiber.Ctx) error {
	todoCollection := database.GetCollection("todos")
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = todoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

func DeleteTodo(c *fiber.Ctx) error {
	todoCollection := database.GetCollection("todos")
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	_, err = todoCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
