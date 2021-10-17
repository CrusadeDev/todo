package query

import (
	"todoApp/internal/models"
	"todoApp/internal/repository"
)

type IndexHandler struct {
	todoRepository repository.TaskRepositoryInterface
}

func NewIndexHandler(todoRepository repository.TaskRepositoryInterface) *IndexHandler {
	return &IndexHandler{todoRepository: todoRepository}
}

func (h *IndexHandler) Handle() []models.Task {
	return h.todoRepository.GetAll()
}
