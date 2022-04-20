package routes

import (
	"govirt/configs"
	"govirt/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, configuration *configs.Config) {
	e.GET("/ping", controllers.Ping(configuration))
	e.POST("/oauth/token", controllers.RequestToken(configuration))
}