package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("error loading .env file")
  }

  DB_URL := os.Getenv("DB_URL")
  clientOption := options.Client().ApplyURI(DB_URL)
  client, err := mongo.Connect(clientOption)

  if err != nil {
    log.Fatal(err)
  }

  err = client.Ping(context.Background(), nil)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("connected to db")

}