package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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

	http.Handle("/", taskController)
	http.ListenAndServe(":8080", nil)
}

func initDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id VARCHAR(50) NOT NULL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT
	)`
	_, err := db.Exec(query)
	return err
}
