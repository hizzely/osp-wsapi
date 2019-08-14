package main

import (
	Routes "github.com/hizzely/osp-wsapi/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	echo := echo.New()
	echo.HideBanner = true

	// Middleware
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	// Routes
	Routes.Register(echo)

	// Start server
	echo.Logger.Fatal(echo.Start(":8200"))
}
