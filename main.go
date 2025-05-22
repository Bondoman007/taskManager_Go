package main

import (
	"fmt"
	"net/http"

	"github.com/Bondoman007/taskManager_Go/routes"
)

func main() {
	fmt.Println("Starting Task Service on :8080...")
	http.ListenAndServe(":8080", routes.SetupRouter())
}
