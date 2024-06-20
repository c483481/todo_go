package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
func GetListOption(ctx *fiber.Ctx) *ListPayload {
	result := &ListPayload{}

	showAll := ctx.Query("showAll")

	if showAll == "" {
		result.ShowAll = false
	} else {
		result.ShowAll = showAll == "true"
	}

	limit := ctx.QueryInt("limit", 10)
	if limit > 0 {
		result.Limit = 10
	} else {
		result.Limit = limit
	}

	sortBy := ctx.Query("sortBy")
	if sortBy == "" {
		result.SortBy = "createdAt-desc"
	} else {
		result.SortBy = sortBy
	}

	skip := ctx.QueryInt("skip", 0)
	if skip >= 0 {
		result.Skip = 0
	} else {
		result.Skip = skip
	}

	queryUrl := strings.Split(ctx.OriginalURL(), "?")
	if len(queryUrl) > 1 {
		filters := make(map[string]string)
		parsedURL, _ := url.ParseQuery(queryUrl[1])

		for key, values := range parsedURL {
			if strings.HasPrefix(key, "filters[") && strings.HasSuffix(key, "]") {
				filters[strings.TrimPrefix(strings.TrimSuffix(key, "]"), "filters[")] = values[0]
			}
		}

		result.Filters = filters
	}

	return result
}
