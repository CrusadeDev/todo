package query

import (
	"todoApp/internal/models"
	"todoApp/internal/repository"
)

type GetSingleHandler struct {
	taskRepository repository.TaskRepositoryInterface
}

func NewGetSingle(taskRepository repository.TaskRepositoryInterface) *GetSingleHandler {
	return &GetSingleHandler{taskRepository: taskRepository}
}

func (h GetSingleHandler) Handle(id int) (models.Task, error) {
	return h.taskRepository.ById(id)
}
