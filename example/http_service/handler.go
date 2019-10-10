package main

import (
	"net/http"

	"github.com/ManiMuridi/goclean/command"
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
				result := command.Execute(&GetAll{})
				return ctx.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Get User By Name",
			Path:   "/users/:name",
			Method: http.MethodGet,
			Handler: func(ctx *httpservice.Context) error {
				name := ctx.Param("name")
				result := command.Execute(&GetByName{name})
				return ctx.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Update User By Name",
			Path:   "/users/:name",
			Method: http.MethodPut,
			Handler: func(ctx *httpservice.Context) error {
				name := ctx.Param("name")
				//req := &UpdateByNameRequest{Name: name}
				cmd := &Update{UserName: name}
				return ctx.BindableJSONResult(cmd)
				//if err := ctx.Bind(&req.User); err != nil {
				//	return ctx.JSON(http.StatusInternalServerError, err)
				//}
				//
				//result := command.Execute(&Update{req})
				//
				//return ctx.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Create User",
			Path:   "/users",
			Method: http.MethodPost,
			Handler: func(ctx *httpservice.Context) error {
				return ctx.BindableJSONResult(&Create{})
			},
		},
	}
}
