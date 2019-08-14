package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeIndex handler
func HomeIndex(c echo.Context) error {
	return c.String(http.StatusOK, "It works!")
}
