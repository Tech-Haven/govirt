package controllers

import (
	"govirt/configs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(config *configs.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong!")
}
}