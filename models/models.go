package models

import "time"

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
}

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreateAt    time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
