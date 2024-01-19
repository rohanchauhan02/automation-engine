package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rohanchauhan02/automation-engine/domain/reminder"
	"github.com/rohanchauhan02/automation-engine/models"
)

type handlerReminder struct {
	usecase reminder.Usecase
}

func NewHandlerReminder(r chi.Router, usecase reminder.Usecase) {
	handler := &handlerReminder{
		usecase: usecase,
	}
	r.Post("/api/users/{userID}/task", handler.CreateTask)
}

// CreateTask handles the creation of a new task for a user
func (h *handlerReminder) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Parse the userID from the URL parameters
	userID := chi.URLParam(r, "userID")

	// Parse the request body into a Task struct
	payload := models.TaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	createdTask, err := h.usecase.CreateTask(userID, payload)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	
	fmt.Printf("[SUCCESS][CreateTask][Response] success to create task for user: %s", userID)
	// Encode the created task as JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}
