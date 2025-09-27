package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
  Id int `json:"id"`
  IsCompleted bool `json:"isCompleted"`
  Body string `json:"body"`
}

func main() {
	fmt.Println("Hello world air")

  app := fiber.New()

  todos := []Todo{}

  app.Get("/", func(c fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello world"})
  })

  app.Post("/api/todos", func(c fiber.Ctx) error {
    // get address of Todo struct litreal into 'todo' variable
    todo := new(Todo) 

    // parse incoming body request and bind value to creted todo struct litreal i.e. todo (with error handle)
    if err := c.Bind().Body(todo); err != nil {
      return err
    }

    // handle error when body is empty
    if todo.Body == "" {
      return c.Status(400).JSON(fiber.Map{"error": "Todo body should not be empty"})
    }
    // update todo id
    todo.Id = len(todos) + 1
    // append value of todo into 
    todos = append(todos, *todo)

    return c.Status(201).JSON(todos)

  })


  log.Fatal(app.Listen(":3000"))

}