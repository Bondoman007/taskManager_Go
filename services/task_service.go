package services

import (
	"strconv"
	"strings"

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


func GetAllTasks(filterStatus, pageStr, limitStr string) []models.Task {
	all := []models.Task{}

	// Collect matching tasks
	for _, task := range db.Tasks {
		if filterStatus == "" || strings.EqualFold(task.Status, filterStatus) {
			all = append(all, task)
		}
	}

	// Apply pagination
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(all) {
		return []models.Task{}
	}
	if end > len(all) {
		end = len(all)
	}

	return all[start:end]
}
