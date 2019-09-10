package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	result := (&GetAll{}).Execute()
	result.Error = errors.New("something went wrong")
	return c.JSON(http.StatusOK, result)
}
