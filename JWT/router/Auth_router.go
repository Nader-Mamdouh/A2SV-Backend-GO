package router

import (
	"JWT/controllers"
	"JWT/data"
	"JWT/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRouterSetup(r *gin.Engine, userCollection *mongo.Collection) {
	userService := data.NewUserService(userCollection)
	authController := controllers.NewAuthController(userService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			// Add admin-specific routes here
		}

		// User routes
		user := protected.Group("/user")
		user.Use(middleware.RoleMiddleware("user"))
		{
			// Add user-specific routes here
		}
	}
}
