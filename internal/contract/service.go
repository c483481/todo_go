package contract

import (
	"github.com/c483481/todo_go/internal/dto/todos"
	"github.com/c483481/todo_go/pkg/handler"
)

type Service struct {
	Todos TodosService
}

type TodosService interface {
	Create(payload *todos.Payload) (*todos.Result, error)
	Detail(xid string) (*todos.Result, error)
	List(payload *handler.ListPayload) (*handler.FindResult[*todos.Result], error)
	Update(payload *todos.UpdatePayload) (*todos.Result, error)
}
