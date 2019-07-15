package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/tasks", tasksHandler)
	http.ListenAndServe(":8080", nil)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, http!")
}
