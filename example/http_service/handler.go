package main

import (
	"errors"
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
			HandlerFunc: func(c echo.Context) error {
				result := (&GetAll{}).Execute()
				result.Error = errors.New("something went wrong")
				return c.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Get User By Name",
			Path:   "/users/:name",
			Method: http.MethodGet,
			HandlerFunc: func(c echo.Context) error {
				name := c.Param("name")
				result := command.Execute(&GetByName{name})
				return c.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Update User By Name",
			Path:   "/users/:name",
			Method: http.MethodPut,
			HandlerFunc: func(c echo.Context) error {
				name := c.Param("name")
				req := &UpdateByNameRequest{Name: name}

				if err := c.Bind(&req.User); err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}

				result := command.Execute(&Update{req})

				return c.JSON(http.StatusOK, result)
			},
		},
		{
			Name:   "Create User",
			Path:   "/users",
			Method: http.MethodPost,
			HandlerFunc: func(c echo.Context) error {
				req := &CreateRequest{}

				if err := c.Bind(req.User); err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}

				//if err := c.Validate(req.User); err != nil {
				//	return c.JSON(http.StatusBadRequest, err)
				//}

				result := command.Execute(&Create{req})

				return c.JSON(http.StatusOK, result)
			},
		},
	}
}
