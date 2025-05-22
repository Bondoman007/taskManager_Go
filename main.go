package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bondoman007/taskManager_Go/routes"
)

func main() {
    fmt.Println("Starting Task Service on :8080...")

    router := routes.SetupRouter()
    
    loggedRouter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        router.ServeHTTP(w, r)
    })

    log.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", loggedRouter); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}