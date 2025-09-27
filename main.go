package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/lucsky/cuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
  ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
  IsCompleted bool   `json:"isCompleted" bson:"isCompleted"`
  Body        string `json:"body" bson:"body"`
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

  app.Get("/api/todos", getTodo)
  app.Post("/api/todos", addTodo)
  app.Patch("/api/todos/:id", updateTodo)
  // app.Delete("/api/todos", deleteTodo)

  port := os.Getenv("PORT")

  log.Fatal(app.Listen(":"+port))
}

func addTodo(c fiber.Ctx) error {
  // create blueprint of type Todo
    todo := new(Todo)

    // bind value comes from user
    if err := c.Bind().Body(todo); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "failed to parse request body",
        })
    }

     // Always assign a cuid for new todos
    todo.ID = cuid.New()

    result, err := coll.InsertOne(context.TODO(), todo)
    if err != nil {
        log.Printf("Failed to insert todo: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "failed to insert todo into database",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Todo created successfully",
        "todoId": result.InsertedID,
    })
}

func getTodo(c fiber.Ctx) error {
  var todos []Todo

  // Find all documents with an empty filter
  cursor, err := coll.Find(context.TODO(), bson.M{})
  if err != nil {
    // Handle the error gracefully, instead of panicking
    return c.Status(500).JSON(fiber.Map{
      "error": "Failed to find todos",
    })
  }
  defer cursor.Close(context.TODO())

  // Decode all documents from the cursor into the 'todos' slice
  if err = cursor.All(context.TODO(), &todos); err != nil {
    panic(err)
  }

  // Fiber automatically marshals the Go struct slice to a JSON array
  return c.JSON(todos)
}

func updateTodo(c fiber.Ctx) error {
  id := c.Params("id")
  updateDoc := bson.M{
		"$set": bson.M{"isCompleted": true},
	}
  coll.UpdateByID(context.TODO(), id, updateDoc)

  return c.SendStatus(200)
}

