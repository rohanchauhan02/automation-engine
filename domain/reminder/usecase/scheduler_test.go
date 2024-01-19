package usecase

import (
	"testing"

	"github.com/rohanchauhan02/automation-engine/domain/reminder"
)

func Test_usecase_Reminder(t *testing.T) {
	type fields struct {
		repository reminder.Repository
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repository: tt.fields.repository,
			}
			u.Reminder()
		})
	}
}
