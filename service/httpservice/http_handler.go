package httpservice

import (
	"github.com/labstack/echo/v4"
)

type HttpHandler interface {
	Middleware() []echo.MiddlewareFunc
	Routes() []Route
}

type Request struct {
	Context echo.Context
}

type Route struct {
	Method     string
	Path       string
	Name       string
	Handler    func(request *Request) error
	Middleware []echo.MiddlewareFunc
}

type BodyDecoder interface {
	DecodeBody(obj interface{}) error
}
