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
