package usecase

import (
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/automation-engine/domain/mocks/mock_reminder"
	"github.com/rohanchauhan02/automation-engine/models"
	"github.com/rohanchauhan02/automation-engine/util"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

const (
	Assertation = "Assertations"
	Conf        = "shared.config"
	Failed      = "Should have failed"
	Succeeded   = "Should have succeeded"
	Separator   = " - "
)

type TestReminder struct {
	Task []models.Task `json:"task"`
	User []models.User `json:"user"`
}

func Test_usecase_CreateTask(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database: %v", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("Failed to open GORM database: %v", err)
	}
	defer gdb.Close()

	reminderTestData := TestReminder{}
	util.ReadJSON("../../../test_files/reminder_test_data.json", &reminderTestData)
	type args struct {
		userIDStr string
		payload   models.TaskRequest
	}
	tests := []struct {
		name     string
		prepFunc func(mockReminderRepo *mock_reminder.MockRepository)
		args     args
		want     *models.Task
		wantErr  error
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1: Valid reminder",
			prepFunc: func(mockReminderRepo *mock_reminder.MockRepository) {
				task := reminderTestData.Task[0]
				user := reminderTestData.User[0]
				mockReminderRepo.EXPECT().InsertTask(gomock.Any()).AnyTimes().Return(gomock.Any())
			},
			args: args{
				userIDStr: "1",
				payload: struct {
					Title       string `json:"title"`
					Description string `json:"description"`
					Priority    string `json:"priority"`
					DueDate     string `json:"due_date"`
				}{
					Title:       "",
					Description: "",
					Priority:    "",
					DueDate:     "",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockReminderRepo := mock_reminder.NewMockRepository(mockCtrl)
			u := usecase{
				repository: mockReminderRepo,
			}
			if tt.prepFunc != nil {
				tt.prepFunc(mockReminderRepo)
			}
			Convey(tt.name, t, func() {
				_, err := u.CreateTask(tt.args.userIDStr, tt.args.payload)
				//====ASSERTION========
				Convey(strings.Join([]string{tt.name, Assertation}, Separator), func() {
					Convey(strings.Join([]string{tt.name, Succeeded}, Separator), func() {
						So(err, ShouldResemble, tt.wantErr)
					})
					if tt.wantErr != nil || err != nil {
						Convey(strings.Join([]string{tt.name, Failed}, Separator), func() {
							So(err.Error(), ShouldEqual, tt.wantErr.Error())
						})
					}
				})
			})
		})
	}
}
