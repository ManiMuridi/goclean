package main

import (
	"net/http"

	"github.com/ManiMuridi/goclean/command"

	"github.com/ManiMuridi/goclean/service/httpservice"

	"github.com/labstack/echo/v4"
)

type handler struct{}

func (h *handler) Middleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{Cors}
}

func (h *handler) Routes() []httpservice.Route {
	return []httpservice.Route{
		{
			Name:   "Get All Users",
			Path:   "/users",
			Method: http.MethodGet,
			Handler: func(request *httpservice.Request) error {
				result := command.Execute(&GetAll{})
				return request.Context.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Get User By Name",
			Path:   "/users/:name",
			Method: http.MethodGet,
			Handler: func(request *httpservice.Request) error {
				name := request.Context.Param("name")
				result := command.Execute(&GetByName{name})
				return request.Context.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Update User By Name",
			Path:   "/users/:name",
			Method: http.MethodPut,
			Handler: func(request *httpservice.Request) error {
				name := request.Context.Param("name")
				req := &UpdateByNameRequest{Name: name}

				if err := request.Context.Bind(&req.User); err != nil {
					return request.Context.JSON(http.StatusInternalServerError, err)
				}

				result := command.Execute(&Update{req})

				return request.Context.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Create User",
			Path:   "/users",
			Method: http.MethodPost,
			Handler: func(request *httpservice.Request) error {
				req := &CreateRequest{}

				if err := request.Context.Bind(&req.User); err != nil {
					return request.Context.JSON(http.StatusInternalServerError, err)
				}

				result := command.Execute(&Create{req})

				return request.Context.JSON(http.StatusOK, result)
			},
		},
	}
}
