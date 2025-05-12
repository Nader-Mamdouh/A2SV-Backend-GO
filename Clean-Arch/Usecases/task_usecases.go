package usecases

import (
	domain "JWT/Domain"
	repositories "JWT/Repositories"
)

// TaskUsecase handles task-related business logic.
type TaskUsecase struct {
	taskRepo repositories.TaskRepository
}

// NewTaskUsecase creates a new TaskUsecase instance.
func NewTaskUsecase(taskRepo repositories.TaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: taskRepo}
}

// GetTasks retrieves all tasks.
func (t *TaskUsecase) GetTasks() ([]domain.Task, error) {
	return t.taskRepo.GetTasks()
}

// GetTask retrieves a task by ID.
func (t *TaskUsecase) GetTask(taskID string) (domain.Task, error) {
	return t.taskRepo.GetTask(taskID)
}

// UpdateTask updates a task.
func (t *TaskUsecase) UpdateTask(taskID string, task domain.Task) (int64, int64, error) {
	return t.taskRepo.UpdateTask(taskID, task)
}

// RemoveTask removes a task.
func (t *TaskUsecase) RemoveTask(taskID string) (int64, error) {
	return t.taskRepo.RemoveTask(taskID)
}

// AddTask adds a new task.
func (t *TaskUsecase) AddTask(task domain.Task) (interface{}, error) {
	return t.taskRepo.AddTask(task)
}
