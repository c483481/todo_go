package handler

type AppManifest struct {
	AppName    string
	AppVersion string
}

type AppResponses struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Data    any    `json:"data"`
}

type ErrorResponseType struct {
	Status  int
	Code    string
	Message string
	Data    any
}

type ErrorType struct {
	Status  int
	Message string
	Data    any
}

type errorResponseType struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
