package router

import (
	"TaskManagementRESTAPI_MongoDB/controllers"

	"github.com/gin-gonic/gin"
)

// RouterSetup initializes and returns a new Gin router
func RouterSetup() *gin.Engine {
	r := gin.Default()

    r.POST("/tasks", controllers.CreateTask)
    r.GET("/tasks", controllers.GetTasks)
    r.GET("/tasks/:id", controllers.GetTask)
    r.PUT("/tasks/:id", controllers.UpdateTask)
    r.DELETE("/tasks/:id", controllers.DeleteTask)

    return r
}
