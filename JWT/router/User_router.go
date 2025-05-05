package router

import (
	controller "JWT/controllers"
	middleware "JWT/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouterSetup(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:user_id", controller.GetUser)
}
