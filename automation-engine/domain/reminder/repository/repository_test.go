package repository

import (
	"database/sql"
	"testing"
	"time"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rohanchauhan02/automation-engine/models"
	"github.com/stretchr/testify/assert"
)

func TestGetTaskByDueDate(t *testing.T) {
	// Mock the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	// Create a new repository with the mocked database connection
	r := &Repository{db: db}

	// Mock the SQL query execution
	rows := sqlmock.NewRows([]string{"id", "name", "due_date"})
	mock.ExpectQuery("SELECT \* FROM tasks WHERE due_date <= ?").WillReturnRows(rows)

	// Call the function
	tasks, err := r.GetTaskByDueDate(60)

	// Assert that the function returned no error
	assert.NoError(t, err)

	// Assert that the function returned an empty slice of tasks
	assert.Equal(t, []models.Task{}, tasks)

	// Assert that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}