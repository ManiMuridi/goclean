package httpservice

import (
	"net/http"
	"reflect"

	"github.com/ManiMuridi/goclean/syserr"

	"github.com/ManiMuridi/goclean/command"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func (ctx *Context) BindableJSONResult(cmd command.Command, data interface{}) error {
	if reflect.ValueOf(data).Type().Kind() != reflect.Ptr {
		return ctx.JSON(http.StatusInternalServerError, NewResponse(&syserr.ValidationError{"Data": "Data must be a pointer struct"}, nil))
	}

	if data != nil {
		if err := ctx.Bind(data); err != nil {
			return ctx.JSON(http.StatusBadRequest, NewResponse(err, nil))
		}
	} else {
		if err := ctx.Bind(&cmd); err != nil {
			return ctx.JSON(http.StatusBadRequest, NewResponse(err, nil))
		}
	}

	return ctx.JSONResult(cmd)
}

func (ctx *Context) JSONResult(cmd command.Command) error {
	result := command.Execute(cmd)

	switch syserr.Type(result.Error) {
	case "":
		return ctx.JSON(http.StatusOK, NewResponseResult(result))
	case syserr.Validation:
		return ctx.JSON(http.StatusBadRequest, NewResponseResult(result))
	default:
		return ctx.JSON(http.StatusInternalServerError, NewResponseResult(result))
	}
}
