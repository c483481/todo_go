package handler

type AppErrorResponseMap struct {
	data map[string]*ErrorResponseType
}

var intervalServerError = &ErrorResponseType{
	Code:    "E_ERR_0",
	Status:  500,
	Message: "Interval Server Error",
}

var ErrorResponse = &AppErrorResponseMap{}

func (e *ErrorResponseType) Error() string {
	return e.Code
}

func (e *AppErrorResponseMap) InitError(data map[string]*ErrorType) {
	e.data = map[string]*ErrorResponseType{}
	for key, value := range data {
		e.data[key] = &ErrorResponseType{
			Code:    key,
			Status:  value.Status,
			Message: value.Message,
			Data:    value.Data,
		}
	}
}

func (e *AppErrorResponseMap) GetError(str string) error {
	data, ok := e.data[str]
	if ok {
		return data
	} else {
		return intervalServerError
	}
}

func (e *AppErrorResponseMap) GetIntervalError() error {
	return intervalServerError
}

func (e *AppErrorResponseMap) GetBadRequestError() error {
	return &ErrorResponseType{
		Status:  400,
		Message: "Bad Request",
		Code:    "E_REQ_0",
	}
}

func (e *AppErrorResponseMap) GetUnauthorizedError() error {
	return &ErrorResponseType{
		Status:  401,
		Message: "Unauthorized",
		Code:    "E_AUTH",
	}
}

func (e *AppErrorResponseMap) GetForbiddenError() error {
	return &ErrorResponseType{
		Status:  403,
		Message: "Forbidden",
		Code:    "E_AUTH",
	}
}
