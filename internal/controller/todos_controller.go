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
	app.Get("/:xid", t.GetDetail)
	app.Get("/", t.GetList)
	app.Put("/:xid", t.PutUpdateTodos)
	app.Delete("/:xid", t.DeleteTodos)
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

func (t *todosController) GetDetail(ctx *fiber.Ctx) error {
	xid := ctx.Params("xid")

	result, err := t.service.Detail(xid)

	if err != nil {
		return err
	}

	return handler.WrapData(ctx, result)
}

func (t *todosController) GetList(ctx *fiber.Ctx) error {
	payload := handler.GetListOption(ctx)

	result, err := t.service.List(payload)

	if err != nil {
		return err
	}

	return handler.WrapData(ctx, result)
}

func (t *todosController) PutUpdateTodos(ctx *fiber.Ctx) error {
	payload := &todos.UpdatePayload{}

	err := t.validateBody(ctx, payload)

	if err != nil {
		return err
	}

	payload.Xid = ctx.Params("xid")
	result, err := t.service.Update(payload)

	if err != nil {
		return err
	}

	return handler.WrapData(ctx, result)
}

func (t *todosController) DeleteTodos(ctx *fiber.Ctx) error {
	xid := ctx.Params("xid")

	err := t.service.Delete(xid)

	if err != nil {
		return err
	}

	return handler.WrapData(ctx, "success")
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
