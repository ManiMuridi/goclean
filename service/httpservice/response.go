package httpservice

import "github.com/ManiMuridi/goclean/command"

type Response struct {
	Errors error
	Data   interface{}
}

func NewResponse(errs error, data interface{}) *Response {
	return &Response{Errors: errs, Data: data}
}

func NewResponseResult(result *command.Result) *Response {
	return &Response{
		Errors: result.Error,
		Data:   result.Data,
	}
}
