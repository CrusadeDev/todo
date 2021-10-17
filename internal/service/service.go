package service

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"todoApp/internal/app"
	"todoApp/internal/app/command"
	"todoApp/internal/app/query"
	"todoApp/internal/models"
	"todoApp/internal/repository"
	"todoApp/internal/routes"
)

func NewApplication() app.Application {
	log := logrus.New()

	db, err := gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&models.Task{})

	if err != nil {
		log.Panic(err)
	}

	application := buildApplication(db)

	r := routes.BuildHttpRouter(application, log)
	err = r.Run()

	if err != nil {
		log.Panic(err)
	}

	return application
}

func buildApplication(db *gorm.DB) app.Application {
	todoRepository := repository.NewTodoRepositorySqlite(db)

	return app.Application{
		Queries: app.Queries{
			ShowAll:   *query.NewIndexHandler(todoRepository),
			GetSingle: *query.NewGetSingle(todoRepository),
		},
		Commands: app.Commands{
			Create: *command.NewCreateHandler(todoRepository),
			Remove: *command.NewRemoveHandler(todoRepository),
		},
	}
}
