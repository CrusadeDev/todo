package repository

import (
	"errors"
	"gorm.io/gorm"
	"todoApp/internal/models"
)

type TaskRepositorySqlite struct {
	db *gorm.DB
}

func (t TaskRepositorySqlite) GetAll() []models.Task {
	var item []models.Task
	t.db.Find(&item)

	return item
}

func (t TaskRepositorySqlite) Save(todo models.Task) {
	t.db.Create(todo)
}

func (t TaskRepositorySqlite) Remove(id int) {
	t.db.Delete(&models.Task{}, id)
}

func (t TaskRepositorySqlite) ById(id int) (models.Task, error) {
	var item models.Task
	result := t.db.First(&item, id)

	if result.Error != nil {
		return models.Task{}, errors.New("task not found")
	}

	return item, nil
}

func NewTodoRepositorySqlite(db *gorm.DB) *TaskRepositorySqlite {
	return &TaskRepositorySqlite{db: db}
}
