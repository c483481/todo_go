package controller

import (
	"fmt"
	"github.com/c483481/todo_go/internal/contract"
	"github.com/c483481/todo_go/internal/dto/todos"
	"github.com/c483481/todo_go/pkg/handler"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type todosController struct {
	service  contract.TodosService
	validate *validator.Validate
}

func implTodoController() controller {
	return &todosController{}
}

func (t *todosController) getPrefix() string {
	return "todos"
}

func (t *todosController) initService(service *contract.Service, validate *validator.Validate) {
	t.service = service.Todos
	t.validate = validate
}

func (t *todosController) initRoute(app fiber.Router) {
	app.Post("/", t.PostCreate)
}

func (t *todosController) PostCreate(ctx *fiber.Ctx) error {
	payload := &todos.Payload{}

	err := t.validateBody(ctx, payload)

	if err != nil {
		return err
	}

	result, err := t.service.Create(payload)

	if err != nil {
		return err
	}

	return handler.WrapData(ctx, result)
}

func (t *todosController) validateBody(ctx *fiber.Ctx, data any) error {
	err := ctx.BodyParser(data)

	if err != nil {
		fmt.Println(err)
		return handler.ErrorResponse.GetBadRequestError()
	}

	err = t.validate.Struct(data)

	if err != nil {
		fmt.Println(err)
		return handler.ErrorResponse.GetBadRequestError()
	}

	return nil
}
