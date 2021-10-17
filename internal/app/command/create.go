package command

import (
	"todoApp/internal/models"
	"todoApp/internal/repository"
)

type CreateCommand struct {
	Id      int
	Message string
}

type CreateHandler struct {
	todoRepository repository.TaskRepositoryInterface
}

func NewCreateHandler(todoRepository repository.TaskRepositoryInterface) *CreateHandler {
	return &CreateHandler{todoRepository: todoRepository}
}

func (h *CreateHandler) Handle(C CreateCommand) {
	todo := models.Task{ID: C.Id, Item: C.Message, Completed: false}
	h.todoRepository.Save(todo)
}
