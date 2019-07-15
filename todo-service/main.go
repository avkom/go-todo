package main

import (
	"net/http"
)

func main() {
	taskController := NewTaskController()
	taskRouter := NewTaskRouter(taskController)

	http.Handle("/", taskRouter)
	http.ListenAndServe(":8080", nil)
}
