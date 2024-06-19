package controller

import (
	"github.com/c483481/todo_go/internal/contract"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type controller interface {
	getPrefix() string
	initService(service *contract.Service, validate *validator.Validate)
	initRoute(app fiber.Router)
}
