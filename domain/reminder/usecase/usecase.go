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
		fmt.Printf("[ERROR] failed to parse time. Error: %s", err)
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

	if err != nil {
		fmt.Printf("[ERROR][CreateTask][RESPONSE] failed to insert task. Error: %s", err)
		return nil, err
	}
	return resp, nil
}
