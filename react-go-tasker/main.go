package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	log.Print("Server is running  3000")

	app.Get("/", func (c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Post("/api/todos", func (c fiber.Ctx) error  {
		todo := &Todo{}

		if err := c.Bind().Body(todo); err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"status": fiber.ErrBadRequest,
				"message": "Invalid request body" + err.Error(),
			})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	app.Patch("/api/todos/:id", func (c fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if id == fmt.Sprint(todo.ID) {
				todos[i].Completed = true
				return c.Status(fiber.StatusOK).JSON(todo)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "todo not found",
		})
	})

	app.Listen(":3000")
	
}
