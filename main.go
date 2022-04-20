package main

import (
	"govirt/configs"
	"govirt/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Load .env file
	configs.LoadEnv()

	// Get configurations
	configuration := configs.New()

	// Routes
	routes.Routes(e, configuration)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}