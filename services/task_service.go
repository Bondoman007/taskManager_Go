package services

import (
	"github.com/Bondoman007/taskManager_Go/db"
	"github.com/Bondoman007/taskManager_Go/models"
)

// CreateTask adds a task to the in-memory DB
func CreateTask(task models.Task) {
	db.Tasks[task.ID] = task
}

// GetTask retrieves a task by ID
func GetTask(id string) (models.Task, bool) {
	task, exists := db.Tasks[id]
	return task, exists
}

// UpdateTask modifies an existing task
func UpdateTask(id string, updated models.Task) bool {
	if _, exists := db.Tasks[id]; !exists {
		return false
	}
	db.Tasks[id] = updated
	return true
}

// DeleteTask removes a task
func DeleteTask(id string) bool {
	if _, exists := db.Tasks[id]; !exists {
		return false
	}
	delete(db.Tasks, id)
	return true
}
