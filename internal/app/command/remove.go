package command

import (
	"todoApp/internal/repository"
)

type RemoveCommand struct {
	Id int
}

type RemoveHandler struct {
	TodoRepository repository.TaskRepositoryInterface
}

func NewRemoveHandler(todoRepository repository.TaskRepositoryInterface) *RemoveHandler {
	return &RemoveHandler{TodoRepository: todoRepository}
}

func (h RemoveHandler) Handle(c RemoveCommand) {
	h.TodoRepository.Remove(c.Id)
}
