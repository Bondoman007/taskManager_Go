package routes

import (
	"net/http"

	"github.com/Bondoman007/taskManager_Go/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Task Manager API!"))
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handlers.GetTasks)
		r.Post("/", handlers.CreateTask)
		r.Get("/{id}", handlers.GetTask)
		r.Put("/{id}", handlers.UpdateTask)
		r.Delete("/{id}", handlers.DeleteTask)
		r.Get("/status/{status}", handlers.GetTasksByStatus)
	})

	return r
}
