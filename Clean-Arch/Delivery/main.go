package main

import (
	"context"
	"log"
	"time"

	routers "JWT/Delivery/routers"
	infrastructure "JWT/Infrastructure"
	repositories "JWT/Repositories"
	usecases "JWT/Usecases"

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

	// Initialize services
	jwtService := infrastructure.NewJWTService()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(userCollection)

	// Initialize usecases
	userUsecase := usecases.NewUserUsecase(userRepo)

	// Initialize router
	r := routers.SetupRouter(userUsecase, jwtService)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
