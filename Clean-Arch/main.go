package main

import (
	"context"
	"log"
	"time"

	routes "JWT/router"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Initialize database and collection
	database := client.Database("task_manager")
	userCollection := database.Collection("users")

	// Initialize Gin router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set up routes
	routes.AuthRouterSetup(r, userCollection)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
