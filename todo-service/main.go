package main

import (
	"net/http"
)

func main() {
	taskController := NewTaskController()

	http.Handle("/", taskController)
	http.ListenAndServe(":8080", nil)
}
