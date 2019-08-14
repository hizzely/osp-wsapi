package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
 * ------------------------------------------------------------------
 * Dosen / Account
 * ------------------------------------------------------------------
 */

// DosenAccountIndex handler
func DosenAccountIndex(c echo.Context) error {
	return c.String(http.StatusOK, "DosenAccountIndex")
}

// DosenAccountStore handler
func DosenAccountStore(c echo.Context) error {
	return c.String(http.StatusOK, "DosenAccountStore")
}

// DosenAccountDelete handler
func DosenAccountDelete(c echo.Context) error {
	return c.String(http.StatusOK, "DosenAccountDelete. params id: "+c.Param("id"))
}

// DosenAccountLogin handler
func DosenAccountLogin(c echo.Context) error {
	return c.String(http.StatusOK, "DosenAccountLogin")
}

/*
 * ------------------------------------------------------------------
 * Dosen / Lecturer
 * ------------------------------------------------------------------
 */

// DosenLectureSubjectDetail handler
func DosenLectureSubjectDetail(c echo.Context) error {
	return c.String(http.StatusOK, "DosenLectureSubjectDetail")
}

// DosenLectureSubjectStore handler
func DosenLectureSubjectStore(c echo.Context) error {
	return c.String(http.StatusOK, "DosenLectureSubjectStore")
}

/*
 * ------------------------------------------------------------------
 * Dosen / Presence
 * ------------------------------------------------------------------
 */

// DosenPresenceSessionDetail handler
func DosenPresenceSessionDetail(c echo.Context) error {
	return c.String(http.StatusOK, "DosenPresenceSessionDetail")
}

// DosenPresenceSessionStore handler
func DosenPresenceSessionStore(c echo.Context) error {
	return c.String(http.StatusOK, "DosenPresenceSessionStore")
}

// DosenPresenceSessionUpdate handler
func DosenPresenceSessionUpdate(c echo.Context) error {
	return c.String(http.StatusOK, "DosenPresenceSessionUpdate")
}

// DosenPresenceSessionDelete handler
func DosenPresenceSessionDelete(c echo.Context) error {
	return c.String(http.StatusOK, "DosenPresenceSessionDelete")
}

// DosenPresenceSessionRefreshCode handler
func DosenPresenceSessionRefreshCode(c echo.Context) error {
	return c.String(http.StatusOK, "DosenPresenceSessionRefreshCode")
}
