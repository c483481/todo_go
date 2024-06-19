package repository

import (
	"github.com/c483481/todo_go/internal/contract"
	"github.com/c483481/todo_go/internal/models"
	"gorm.io/gorm"
)

type todosRepository struct {
	db *gorm.DB
}

func implTodosRepository(db *gorm.DB) contract.TodosRepository {
	return &todosRepository{
		db: db,
	}
}

func (t *todosRepository) Create(todo *models.Todos) error {
	result := t.db.Create(&todo)
	return result.Error
}
