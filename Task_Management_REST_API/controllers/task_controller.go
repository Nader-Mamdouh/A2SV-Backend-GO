package controllers

import (
	"Task_Management_REST_API/data"
	"Task_Management_REST_API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tasks": data.Tasks})
}

func GetTask(c *gin.Context) {
	taskid := c.Param("id")
	id, err := strconv.Atoi(taskid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	for _, task := range data.Tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, gin.H{"task": task})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func UpdateTask(c *gin.Context) {
	taskid := c.Param("id")
	id, err := strconv.Atoi(taskid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	for i, task := range data.Tasks {
		if task.ID == id {
			var updatedTask models.Task
			if err := c.ShouldBindJSON(&updatedTask); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			data.Tasks[i] = updatedTask
			c.JSON(http.StatusOK, gin.H{"task": updatedTask})
			return
		}
	}
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	data.Tasks = append(data.Tasks, newTask)
	c.JSON(http.StatusCreated, gin.H{"task": newTask})
}

func DeleteTask(c *gin.Context) {
	taskid := c.Param("id")
	id, err := strconv.Atoi(taskid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
}
