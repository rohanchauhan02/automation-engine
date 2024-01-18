package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/automation-engine/domain/reminder"
)

type repoHandler struct {
	postgresSession *gorm.DB
}

// NewReminderRepository creates a new instance of the reminder repository.
func NewReminderRepository(postgresSession *gorm.DB) reminder.Repository {
	return &repoHandler{
		postgresSession: postgresSession,
	}
}
