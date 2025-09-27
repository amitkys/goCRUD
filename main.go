package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
  ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
  IsCompleted bool               `json:"isCompleted" bson:"isCompleted"`
  Body        string             `json:"body" bson:"body"`
}

var coll *mongo.Collection

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set in .env")
	}

  client, err := mongo.Connect(options.Client().ApplyURI(dbURL))
  if err != nil { 
    panic(err)  
  }

  defer func () {
    if err := client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

	// Sends a ping to confirm a successful connection
  err = client.Ping(context.TODO(), nil)

  if err != nil {
    panic(err)
  }

  fmt.Println("connected to db")

  coll = client.Database("go_crud").Collection("todos")


  app := fiber.New()

  // app.Get("/api/todos", getTodo)
  app.Post("/api/todos", addTodo)
  // app.Patch("/api/todos/:id", updateTodo)
  // app.Delete("/api/todos", deleteTodo)

  port := os.Getenv("PORT")

  log.Fatal(app.Listen(":"+port))
}

func addTodo(c fiber.Ctx) error {
  todo := new(Todo)

  if err := c.Bind().Body(todo); err != nil {
    c.Status(400).JSON(fiber.Map{
      "message": "failed to parse",
    })
  }

  result, err := coll.InsertOne(context.TODO(), todo)
  if err != nil {
    panic(err)
  }

  return c.Status(200).JSON(result)
}

