package reminder

import (
	"github.com/rohanchauhan02/automation-engine/models"
)

type Usecase interface {
	Reminder()
	CreateTask(userID string, payload models.TaskRequest) (*models.Task, error)
}

type Repository interface {
	InsertTask(dto *models.Task) (*models.Task, error)
	GetUserByID(id int) (*models.User, error)
	GetTaskByDueDate(interval int64) ([]models.Task, error)
}
