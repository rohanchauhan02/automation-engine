package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rohanchauhan02/automation-engine/domain/reminder"
)

type handlerReminder struct {
	usecase reminder.Usecase
}

func NewHandlerReminder(r chi.Router, usecase reminder.Usecase) {
	handler := &handlerReminder{
		usecase: usecase,
	}

	r.Get("/api/task/:taskID", handler.GetTask)
	r.Post("/api/task", handler.CreateTask)
}

func (h *handlerReminder) GetTask(w http.ResponseWriter, r *http.Request) {

}

func (h *handlerReminder) CreateTask(w http.ResponseWriter, r *http.Request) {

}
