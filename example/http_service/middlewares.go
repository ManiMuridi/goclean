package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Running Cors")
		return next(c)
	}
}
