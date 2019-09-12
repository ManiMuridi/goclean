package httpservice

import (
	"github.com/labstack/echo/v4"
)

type HttpHandler interface {
	Middleware() []echo.MiddlewareFunc
	Routes() []Route
}

type Route struct {
	Method      string
	Path        string
	Name        string
	HandlerFunc echo.HandlerFunc
	Middleware  []echo.MiddlewareFunc
}

type BodyDecoder interface {
	DecodeBody(obj interface{}) error
}
