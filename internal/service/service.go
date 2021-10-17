package service

import (
	"fmt"
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
	logger := logrus.New()

	todoRepository := buildTodoRepository(logger)
	application := buildApplication(todoRepository)

	r := routes.BuildHttpRouter(application, logger)
	err := r.Run()

	if err != nil {
		logger.Panic(err)
	}

	return application
}

func buildApplication(todoRepository repository.TaskRepositoryInterface) app.Application {

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

func buildTodoRepository(log *logrus.Logger) repository.TaskRepositoryInterface {
	var todoRepository repository.TaskRepositoryInterface
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})

		if err != nil {
			log.Panic(err)
		}

		err = db.AutoMigrate(&models.Task{})

		if err != nil {
			log.Panic(err)
		}
		todoRepository = repository.NewTodoRepositorySqlite(db)
	case "in-memory":
		todoRepository = repository.NewTaskInMemoryRepository()
	default:
		log.Panic(fmt.Sprintf("Invalid database type: %s", dbType))
	}

	return todoRepository
}
