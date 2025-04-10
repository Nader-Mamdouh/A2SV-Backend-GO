package router

import (
	"Task_Management_REST_API/controllers"

	"github.com/gin-gonic/gin"
)

// RouterSetup initializes and returns a new Gin router
func RouterSetup() *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", controllers.GetTasks)          // Endpoint to get all tasks
	r.GET("/tasks/:id", controllers.GetTask)       // Endpoint to get a specific task by ID
	r.PUT("/tasks/:id", controllers.UpdateTask)    // Endpoint to update a specific task by ID
	r.POST("/tasks", controllers.CreateTask)       // Endpoint to create a new task
	r.DELETE("/tasks/:id", controllers.DeleteTask) // Endpoint to delete a specific task by ID
	return r

}
