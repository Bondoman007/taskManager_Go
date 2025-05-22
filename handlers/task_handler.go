package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/Bondoman007/taskManager_Go/models"
	"github.com/Bondoman007/taskManager_Go/services"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ✅ Validate status
	if task.Status != "Pending" && task.Status != "Completed" {
		http.Error(w, "Invalid status — must be 'Pending' or 'Completed'", http.StatusBadRequest)
		return
	}

	task.ID = uuid.New().String()
	services.CreateTask(task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, found := services.GetTask(id)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updatedTask.ID = id
	ok := services.UpdateTask(id, updatedTask)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ok := services.DeleteTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// ✅ Validate status if provided
	if status != "" && status != "Pending" && status != "Completed" {
		http.Error(w, "Invalid status filter — must be 'Pending' or 'Completed'", http.StatusBadRequest)
		return
	}

	tasks := services.GetAllTasks(status, pageStr, limitStr)
	json.NewEncoder(w).Encode(tasks)
}


func GetTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := chi.URLParam(r, "status")
	if status != "Pending" && status != "Completed" {
		http.Error(w, "Invalid status — must be 'Pending' or 'Completed'", http.StatusBadRequest)
		return
	}
	tasks := services.GetTasksByStatus(status)
	json.NewEncoder(w).Encode(tasks)
}
