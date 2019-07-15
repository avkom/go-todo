package main

import (
	"net/http"
)

func main() {
	taskService := NewTaskService()
	taskController := NewTaskController(taskService)

	http.Handle("/", taskController)
	http.ListenAndServe(":8080", nil)
}
