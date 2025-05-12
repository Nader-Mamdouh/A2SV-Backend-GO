package routers

import (
	controllers "JWT/Delivery/controllers"
	infrastructure "JWT/Infrastructure"
	usecases "JWT/Usecases"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the application routes.
func SetupRouter(userUsecase *usecases.UserUsecase, jwtService *infrastructure.JWTService) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authController := controllers.NewAuthController(userUsecase)

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(infrastructure.AuthMiddleware(jwtService))
	{
		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(infrastructure.RoleMiddleware("admin"))
		{
			// Add admin-specific routes here
		}

		// User routes
		user := protected.Group("/user")
		user.Use(infrastructure.RoleMiddleware("user"))
		{
			user.GET("/users", authController.GetAllUsers)
			user.GET("/users/:id", authController.GetUser)
		}
	}

	return r
}
