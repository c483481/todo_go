package controller

import (
	"fmt"
	"github.com/c483481/todo_go/internal/contract"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ImplController(app *fiber.App, service *contract.Service) {
	allController := []controller{
	}

	validate := validator.New()

	for _, c := range allController {
		c.initService(service, validate)
		group := app.Group(fmt.Sprintf("/%s", c.getPrefix()))
		c.initRoute(group)
		fmt.Printf("initiate route /%s\n", c.getPrefix())
	}
}
