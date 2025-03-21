package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "tasks.db")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Create tasks table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		status TEXT DEFAULT "PENDING"
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create tasks table: %v", err)
	}

	return nil
}

func AddTask(task string) error {
	if task == "" {
		return errors.New("task cannot be empty")
	}

	insertQuery := `INSERT INTO tasks (description) VALUES (?);`
	_, err := db.Exec(insertQuery, task)
	if err != nil {
		return fmt.Errorf("failed to add task: %v", err)
	}

	fmt.Printf("Task added to database: %s\n", task)
	return nil
}

func ListTasks() error {
	rows, err := db.Query("SELECT * FROM tasks;")
	if err != nil {
		return fmt.Errorf("failed to query tasks: %v", err)
	}
	defer rows.Close()

	fmt.Println("-------------------------------------------------")
	fmt.Printf("| %-3s | %-30s | %-8s |\n", "No", "Description", "Status")
	fmt.Println("-------------------------------------------------")

	for rows.Next() {
		var id int
		var description string
		var status string
		err := rows.Scan(&id, &description, &status)
		if err != nil {
			return fmt.Errorf("failed to scan task: %v", err)
		}
		fmt.Printf("| %-3d | %-30s | %-8s |\n", id, description, status)
	}
	fmt.Println("-------------------------------------------------")

	return nil
}

func UpdateTask(id int) error {
	updateQuery := `UPDATE tasks SET status = 'DONE' WHERE id = ?;`
	_, err := db.Exec(updateQuery, id)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	fmt.Printf("Task updated: %d\n", id)
	return nil
}

func DeleteTask(id int) error {
	updateQuery := `DELETE FROM tasks WHERE id = ?;`
	_, err := db.Exec(updateQuery, id)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	fmt.Printf("Task deleted: %d\n", id)
	return nil
}

func CloseDB() error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return fmt.Errorf("failed to close database: %v", err)
		}
	}
	return nil
}
