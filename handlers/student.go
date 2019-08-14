package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
 * ------------------------------------------------------------------
 * Student / Account
 * ------------------------------------------------------------------
 */

// StudentAccountDetail handler
func StudentAccountDetail(c echo.Context) error {
	return c.String(http.StatusOK, "StudentAccountDetail")
}

// StudentAccountStore handler
func StudentAccountStore(c echo.Context) error {
	return c.String(http.StatusOK, "StudentAccountStore")
}

// StudentAccountDelete handler
func StudentAccountDelete(c echo.Context) error {
	return c.String(http.StatusOK, "StudentAccountDelete")
}

// StudentAccountLogin handler
func StudentAccountLogin(c echo.Context) error {
	return c.String(http.StatusOK, "StudentAccountLogin")
}

/*
 * ------------------------------------------------------------------
 * Student / Classroom
 * ------------------------------------------------------------------
 */

// StudentClassroomCoursesIndex handler
func StudentClassroomCoursesIndex(c echo.Context) error {
	return c.String(http.StatusOK, "StudentClassroomCoursesIndex")
}

/*
 * ------------------------------------------------------------------
 * Student / Presence
 * ------------------------------------------------------------------
 */

// StudentPresenceDetail handler
func StudentPresenceDetail(c echo.Context) error {
	return c.String(http.StatusOK, "StudentPresenceDetail")
}

// StudentPresenceHistory handler
func StudentPresenceHistory(c echo.Context) error {
	return c.String(http.StatusOK, "StudentPresenceHistory")
}

// StudentPresenceByCode handler
func StudentPresenceByCode(c echo.Context) error {
	return c.String(http.StatusOK, "StudentPresenceByCode")
}

// StudentPresenceByRfid handler
func StudentPresenceByRfid(c echo.Context) error {
	return c.String(http.StatusOK, "StudentPresenceByRfid")
}
