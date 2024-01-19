package usecase

import (
	"reflect"
	"testing"

	"github.com/rohanchauhan02/automation-engine/domain/reminder"
	"github.com/rohanchauhan02/automation-engine/models"
)

func Test_usecase_CreateTask(t *testing.T) {
	type fields struct {
		repository reminder.Repository
	}
	type args struct {
		userIDStr string
		payload   models.TaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repository: tt.fields.repository,
			}
			got, err := u.CreateTask(tt.args.userIDStr, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
