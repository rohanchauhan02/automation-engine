package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rohanchauhan02/automation-engine/domain/reminder"
	"github.com/rohanchauhan02/automation-engine/models"
)

type usecase struct {
	repository reminder.Repository
}

func NewReminderUsecase(repository reminder.Repository) reminder.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateTask(userIDStr string, payload models.TaskRequest) (*models.Task, error) {

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}

	dueDate, err := time.Parse(time.RFC3339, payload.DueDate)
	if err != nil {
		return nil, err
	}

	task := models.Task{
		Title:       payload.Title,
		Description: payload.Description,
		UserID:      userID,
		DueDate:     dueDate,
		Priority:    payload.Priority,
	}

	resp, err := u.repository.InsertTask(&task)

	fmt.Println(resp, err)

	return nil, nil
}
