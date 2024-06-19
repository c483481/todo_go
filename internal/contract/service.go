package contract

import "github.com/c483481/todo_go/internal/dto/todos"

type Service struct {
	Todos TodosService
}

type TodosService interface {
	Create(payload *todos.Payload) (*todos.Result, error)
}
