package contract

import "github.com/c483481/todo_go/internal/models"

type Repository struct {
	Todos TodosRepository
}

type TodosRepository interface {
	Create(todo *models.Todos) error
}
