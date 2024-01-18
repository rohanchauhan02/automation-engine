package usecase

import "github.com/rohanchauhan02/automation-engine/domain/reminder"

type usecaseHandler struct {
	repository reminder.Repository
}

func NewReminderUsecase(repository reminder.Repository) reminder.Usecase {
	return &usecaseHandler{
		repository: repository,
	}
}
