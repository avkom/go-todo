package main

import (
	"net/http"
)

func main() {
	taskRepository := NewSQLTaskRepository()
	taskService := NewTaskService(taskRepository)
	taskController := NewTaskController(taskService)

	http.Handle("/", taskController)
	http.ListenAndServe(":8080", nil)
}
