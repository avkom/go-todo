package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dataSourceName := os.Getenv("POSTGRES_CONNECTION")
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = initDB(db)
	if err != nil {
		log.Fatal(err)
	}

	taskRepository := NewSQLTaskRepository(db)
	taskService := NewTaskService(taskRepository)
	taskController := NewTaskController(taskService)

	router := mux.NewRouter()
	router.PathPrefix("/api/tasks").Handler(taskController)
	router.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
	router.PathPrefix("/").HandlerFunc(indexHandler)

	http.ListenAndServe(":8080", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexHandler")
	http.ServeFile(w, r, "./index.html")
}

func initDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id VARCHAR(36) NOT NULL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT
	)`
	_, err := db.Exec(query)
	return err
}
