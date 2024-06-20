package contract

import (
	"github.com/c483481/todo_go/internal/models"
	"github.com/c483481/todo_go/pkg/handler"
)

type Repository struct {
	Todos TodosRepository
}

type TodosRepository interface {
	Create(todo *models.Todos) error
	FindByXid(xid string) (*models.Todos, error)
	FindList(payload *handler.ListPayload) (*handler.FindResult[*models.Todos], error)
}
