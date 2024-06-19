package service

import (
	"fmt"
	"github.com/c483481/todo_go/internal/contract"
	"github.com/c483481/todo_go/internal/dto/todos"
	"github.com/c483481/todo_go/internal/models"
	"github.com/c483481/todo_go/pkg/handler"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

type todoService struct {
	todo contract.TodosRepository
}

func implTodosService(repo *contract.Repository) contract.TodosService {
	return &todoService{
		todo: repo.Todos,
	}
}

func (t *todoService) Create(payload *todos.Payload) (*todos.Result, error) {
	item := &models.Todos{
		Xid:         ulid.Make().String(),
		Version:     1,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		Title:       payload.Title,
		Description: payload.Description,
	}

	err := t.todo.Create(item)

	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "dial tcp") {
			return nil, handler.ErrorResponse.GetError("E_CONN_1")
		}
		return nil, handler.ErrorResponse.GetIntervalError()
	}

	return composeTodo(item), nil
}

func composeTodo(row *models.Todos) *todos.Result {
	return &todos.Result{
		Xid:         row.Xid,
		Title:       row.Title,
		Description: row.Description,
		Version:     row.Version,
		UpdatedAt:   row.UpdatedAt.Unix(),
		CreatedAt:   row.CreatedAt.Unix(),
	}
}
