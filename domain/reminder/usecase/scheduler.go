package usecase

import (
	"fmt"
	"log"
	"time"
)

func (u *usecase) Reminder() {
	// Run reminder every 15 minutes
	const interval = 1
	ticker := time.NewTicker(interval * time.Minute)

	for {
		select {
		case <-ticker.C:
			tasks, err := u.repository.GetTaskByDueDate(interval)
			for _, task := range tasks {
				userID := task.UserID
				user, err := u.repository.GetUserByID(userID)
				if err != nil {
					fmt.Printf("[ERROR][REMINDER] failed to get user details. Error: %s", err)
				}
				fmt.Printf("[REMINDER] Mobile no: %s\nName :%s", user.Name, user.Name)
			}

			if err != nil {
				log.Println("Error fetching data from DB:", err)
			}
		}
	}
}
