package main

import (
	"fmt"
	"net/http"

	"github.com/Bondoman007/taskManager_Go/routes"
)

func main() {
	fmt.Println("Starting Task Service on :8080...") // This will print first
	err := http.ListenAndServe(":8080", routes.SetupRouter())
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
