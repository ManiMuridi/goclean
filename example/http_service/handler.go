package main

import (
	"net/http"

	"github.com/ManiMuridi/goclean/service/httpservice"

	"github.com/labstack/echo/v4"
)

type handler struct{}

func (h *handler) Middleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (h *handler) Routes() []httpservice.Route {
	return []httpservice.Route{
		{
			Name:   "Get All Users",
			Path:   "/users",
			Method: http.MethodGet,
			Handler: func(ctx *httpservice.Context) error {
				return ctx.JSONResult(&GetAll{})
			},
		},
		{
			Name:   "Get User By Name",
			Path:   "/users/:name",
			Method: http.MethodGet,
			Handler: func(ctx *httpservice.Context) error {
				return ctx.JSONResult(&GetByName{Name: ctx.Param("name")})
			},
		},
		{
			Name:   "Update User By Name",
			Path:   "/users/:name",
			Method: http.MethodPut,
			Handler: func(ctx *httpservice.Context) error {
				name := ctx.Param("name")
				req := &UpdateByNameRequest{Name: name}
				cmd := &Update{Request: req}
				return ctx.BindableJSONResult(cmd, req)
			},
		},
		{
			Name:   "Create User",
			Path:   "/users",
			Method: http.MethodPost,
			Handler: func(ctx *httpservice.Context) error {
				req := &CreateRequest{}
				return ctx.BindableJSONResult(&Create{Request: req, More: req}, &req.User)
			},
		},
	}
}
