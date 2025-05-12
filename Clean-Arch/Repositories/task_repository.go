package repositories

import domain "JWT/Domain"

// TaskRepository defines the interface for task data access.
type TaskRepository interface {
	GetTasks() ([]domain.Task, error)
	GetTask(taskID string) (domain.Task, error)
	UpdateTask(taskID string, task domain.Task) (int64, int64, error)
	RemoveTask(taskID string) (int64, error)
	AddTask(task domain.Task) (interface{}, error)
}
