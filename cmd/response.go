package main

import "github.com/c483481/todo_go/pkg/handler"

func initErrorResponse() {
	handler.ErrorResponse.InitError(map[string]*handler.ErrorType{
		"E_FOUND_1": {
			Status:  400,
			Message: "Resource Not Found",
		},
		"E_AUTH_1": {
			Status:  401,
			Message: "Unauthorized",
		},
		"E_AUTH_2": {
			Status:  403,
			Message: "Forbidden",
		},
		"E_CONN_1": {
			Status:  503,
			Message: "Database Connection Error",
		},
	})
}
