package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
 * ------------------------------------------------------------------
 * Classroom
 * ------------------------------------------------------------------
 */

// ClassroomIndex handler
func ClassroomIndex(c echo.Context) error {
	return c.String(http.StatusOK, "ClassroomIndex")
}

// ClassroomStore handler
func ClassroomStore(c echo.Context) error {
	return c.String(http.StatusOK, "ClassroomStore")
}
