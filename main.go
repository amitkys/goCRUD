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

  coll = client.Database("crud_go").Collection("todo")

	fmt.Println("Successfully connected to MongoDB")

  app := fiber.New()

  app.Get("/api/todos", getTodo)
  // app.Post("/api/todos", addTodo)
  // app.Patch("/api/todos/:id", updateTodo)
  // app.Delete("/api/todos", deleteTodo)

  port := os.Getenv("PORT")

  log.Fatal(app.Listen(":"+port))
}

func getTodo(c fiber.Ctx) error {
  
}

