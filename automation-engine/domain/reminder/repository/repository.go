package repository

import (
	"database/sql"
	"time"
	"github.com/rohanchauhan02/automation-engine/models"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) GetTaskByDueDate(interval int64) ([]models.Task, error) {
	// Get the current time
	currentTime := time.Now()

	// Calculate the due date based on the interval
	dueDate := currentTime.Add(time.Duration(interval) * time.Minute)

	// Prepare the SQL query
	query := "SELECT * FROM tasks WHERE due_date <= ?"

	// Execute the SQL query
	rows, err := r.db.Query(query, dueDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the returned rows into tasks
	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}