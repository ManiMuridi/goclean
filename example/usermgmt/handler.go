package main

import (
	"fmt"
	"net/http"

	"github.com/ManiMuridi/goclean/service/httpservice"

	"github.com/labstack/echo/v4"
)

type handler struct{}

func Cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Running Cors")
		return next(c)
	}
}

func (h *handler) Middleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{Cors}
}

func (h *handler) Routes() []httpservice.Route {
	return []httpservice.Route{
		{
			Name:        "Get All Users",
			Path:        "/users",
			Method:      http.MethodGet,
			HandlerFunc: GetUsers,
		},
		//{
		//	Name:   "Get User By Name",
		//	Path:   "/users/:name",
		//	Method: http.MethodGet,
		//	HandlerFunc: func(c echo.Context) error {
		//		res, _ := (&GetByName{}).Execute()
		//		return c.JSON(http.StatusOK, res)
		//	},
		//},
		//{
		//	Name:   "Update User By Name",
		//	Path:   "/users/:name",
		//	Method: http.MethodPut,
		//	HandlerFunc: func(c echo.Context) error {
		//		res, _ := (&Update{}).Execute()
		//		return c.JSON(http.StatusOK, res)
		//	},
		//},
		{
			Name:   "Create User",
			Path:   "/users",
			Method: http.MethodPost,
			HandlerFunc: func(c echo.Context) error {
				req := CreateRequest{}

				if err := c.Bind(&req.User); err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}

				//if err := c.Validate(req.User); err != nil {
				//	return c.JSON(http.StatusBadRequest, err)
				//}

				result := (&Create{
					Request: req,
				}).Execute()

				return c.JSON(http.StatusOK, result)
			},
		},
	}
}

//func (h *handler) Response(data interface{}) *goclean.HttpResponse {
//	return &goclean.HttpResponse{Data: data}
//}

//func (h *handler) Initialize(svc goclean.Service) {
//	h.svc = svc
//	svc.Logger().Debug().Msg("Service Initialized")
//}
