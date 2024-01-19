package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/automation-engine/domain/reminder"
	"github.com/rohanchauhan02/automation-engine/models"
)

type repository struct {
	postgresSession *gorm.DB
}

// NewReminderRepository creates a new instance of the reminder repository.
func NewReminderRepository(postgresSession *gorm.DB) reminder.Repository {
	return &repository{
		postgresSession: postgresSession,
	}
}
func (r *repository) InsertTask(dto *models.Task) (*models.Task, error) {
	db := r.postgresSession.Create(dto)
	return dto, db.Error
}

func (r *repository) GetUserByID(id int) (*models.User, error) {
	dto := &models.User{}
	db := r.postgresSession.Where("id = ?", id).Find(&dto)
	return dto, db.Error
}

// Execute the query to retrieve records with due_date within the T-minute interval
func (r *repository) GetTaskByDueDate(interval int64) ([]models.Task, error) {
	dto := []models.Task{}
	// Calculate the time threshold for the T-minute interval
	threshold := time.Now().Add(-time.Duration(interval) * time.Minute)
	db := r.postgresSession.Where("due_date BETWEEN ? AND ?", threshold, time.Now()).Find(&dto)
	return dto, db.Error
}
