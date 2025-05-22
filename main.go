package main

import (
	"fmt"
	"net/http"

	"github.com/Bondoman007/taskManager_Go/routes"
)

func main() {

	http.ListenAndServe(":8080", routes.SetupRouter())
	fmt.Println("Starting Task Service on :8080...")

}
