package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleApiStatus(manifest *AppManifest) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(manifest)
	}
}

func HandleNotFound() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(&AppResponses{
			Success: false,
			Code:    "Not Found",
			Data:    fmt.Sprintf("Not Found path %s", ctx.Path()),
		})
	}
}

func WrapData(ctx *fiber.Ctx, data interface{}) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return ctx.JSON(&AppResponses{
		Success: true,
		Code:    "OK",
		Data:    data,
	})
}

func HandleError(c *fiber.Ctx, err error) error {
	var e *ErrorResponseType

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if errors.As(err, &e) {
		return c.Status(e.Status).JSON(&errorResponseType{
			Success: false,
			Code:    e.Code,
			Message: e.Message,
			Data:    e.Data,
		})
	}

	return c.Status(intervalServerError.Status).JSON(&errorResponseType{
		Success: false,
		Code:    intervalServerError.Code,
		Message: intervalServerError.Message,
		Data:    intervalServerError.Data,
	})
}
