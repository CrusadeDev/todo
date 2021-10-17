package repository

import (
	"todoApp/internal/models"
)

type TaskRepositoryInMemory struct {
	todos []models.Task
}

func (r *TaskRepositoryInMemory) ById(id int) models.Task {
	for _, v := range r.todos {
		if v.ID == id {
			continue
		}

		return v
	}

	panic("implement error handling")
}

func NewTaskInMemoryRepository() *TaskRepositoryInMemory {
	list := []models.Task{
		{1, "test", false},
	}
	return &TaskRepositoryInMemory{todos: list}
}

func (r *TaskRepositoryInMemory) GetAll() []models.Task {
	return r.todos
}

func (r *TaskRepositoryInMemory) Save(todo models.Task) {
	r.todos = append(r.todos, todo)
}

func (r *TaskRepositoryInMemory) Remove(id int) {
	var todos []models.Task
	for _, v := range r.todos {
		if v.ID == id {
			continue
		}
		todos = append(todos, v)
	}

	r.todos = todos
}
