package repository

import (
	"todoApp/internal/models"
)

type TaskRepositoryInterface interface {
	GetAll() []models.Task
	Save(todo models.Task)
	Remove(id int)
	ById(id int) models.Task
}
