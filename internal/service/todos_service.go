package service

import (
	"errors"
	"strings"
	"time"

	"github.com/c483481/todo_go/internal/contract"
	"github.com/c483481/todo_go/internal/dto/todos"
	"github.com/c483481/todo_go/internal/models"
	"github.com/c483481/todo_go/pkg/handler"
	"gorm.io/gorm"

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
		// check if err is database connection loss
		if strings.Contains(err.Error(), "dial tcp") {
			return nil, handler.ErrorResponse.GetError("E_CONN_1")
		}
		return nil, handler.ErrorResponse.GetIntervalError()
	}

	return composeTodo(item), nil
}

func (t *todoService) Detail(xid string) (*todos.Result, error) {
	// parse string xid to ulid
	_, err := ulid.ParseStrict(xid)

	if err != nil {
		// return not found because xid must be ulid
		return nil, handler.ErrorResponse.GetError("E_FOUND_1")
	}

	todo, err := t.todo.FindByXid(xid)

	if err != nil {
		// check if the error is Record Not Found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handler.ErrorResponse.GetError("E_FOUND_1")
		}
		// check if err is database connection loss
		if strings.Contains(err.Error(), "dial tcp") {
			return nil, handler.ErrorResponse.GetError("E_CONN_1")
		}
		return nil, handler.ErrorResponse.GetIntervalError()
	}

	return composeTodo(todo), nil
}

func (t *todoService) List(payload *handler.ListPayload) (*handler.FindResult[*todos.Result], error) {
	items, err := t.todo.FindList(payload)

	if err != nil {
		// check if err is database connection loss
		if strings.Contains(err.Error(), "dial tcp") {
			return nil, handler.ErrorResponse.GetError("E_CONN_1")
		}
		return nil, handler.ErrorResponse.GetIntervalError()
	}

	result := make([]*todos.Result, 0, len(items.Result))

	// compose list result
	for _, item := range items.Result {
		result = append(result, composeTodo(item))
	}

	return &handler.FindResult[*todos.Result]{
		Result: result,
		Count:  items.Count,
	}, nil
}

func (t *todoService) Update(payload *todos.UpdatePayload) (*todos.Result, error) {
	_, err := ulid.ParseStrict(payload.Xid)

	if err != nil {
		// return not found because xid must be ulid
		return nil, handler.ErrorResponse.GetError("E_FOUND_1")
	}

	todo, err := t.todo.FindByXid(payload.Xid)

	if err != nil {
		// check if the error is Record Not Found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handler.ErrorResponse.GetError("E_FOUND_1")
		}
		// check if err is database connection loss
		if strings.Contains(err.Error(), "dial tcp") {
			return nil, handler.ErrorResponse.GetError("E_CONN_1")
		}
		return nil, handler.ErrorResponse.GetIntervalError()
	}

	todo.Version += 1
	todo.Title = payload.Title
	todo.Description = payload.Description

	result, err := t.todo.Update(todo.ID, todo, payload.Version)

	if err != nil {
		// check if err is database connection loss
		if strings.Contains(err.Error(), "dial tcp") {
			return nil, handler.ErrorResponse.GetError("E_CONN_1")
		}
		return nil, handler.ErrorResponse.GetIntervalError()
	}

	// check row affected
	if result <= 0 {
		// throw invalid version if row affected less than or equal 0
		return nil, handler.ErrorResponse.GetError("E_REQ_2")
	}

	return composeTodo(todo), nil
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
