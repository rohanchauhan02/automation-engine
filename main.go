package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	reminderHandler "github.com/rohanchauhan02/automation-engine/domain/reminder/delivery/http"
	"github.com/rohanchauhan02/automation-engine/domain/reminder/repository"
	usecase "github.com/rohanchauhan02/automation-engine/domain/reminder/usecase"
	"github.com/rohanchauhan02/automation-engine/shared/config"
	"github.com/rohanchauhan02/automation-engine/shared/container"
	"github.com/rohanchauhan02/automation-engine/shared/postgres"
	"github.com/rohanchauhan02/automation-engine/util"
)

func main() {
	r := chi.NewRouter()
	container := container.DefaultContainer()
	conf := container.MustGet("shared.config").(config.ImmutableConfigInterface)
	postgress := container.MustGet("shared.database").(postgres.PostgresInterface)
	postgresSession, err := postgress.OpenPostgresConn()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open postgres connection: %s", err.Error())
		fmt.Println(msgError)
		panic(err)
	}
	defer postgresSession.Close()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(util.CustomApplicationContext{
				Context:         r.Context(),
				Container:       container,
				SharedConfig:    conf,
				PostgresSession: postgresSession,
			}))
		})
	})

	reminderRepo := repository.NewReminderRepository(postgresSession)
	reminderUsecase := usecase.NewReminderUsecase(reminderRepo)
	reminderHandler.NewHandlerReminder(r, reminderUsecase)
	go reminderUsecase.Reminder()
	port := conf.GetPort()
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Printf("Server error: %s", err.Error())
	}

}
