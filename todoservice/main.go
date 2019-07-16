package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Print("Starting application")
	godotenv.Load()

	dataSourceName := os.Getenv("POSTGRES_CONNECTION")
	log.Printf("POSTGRES_CONNECTION: %v", dataSourceName)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
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
