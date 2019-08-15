package handlers

import (
	"log"
	"net/http"

	DB "github.com/hizzely/osp-wsapi/database"
	"github.com/hizzely/osp-wsapi/helpers"
	"github.com/labstack/echo/v4"
)

/*
 * ------------------------------------------------------------------
 * Dosen / Account
 * ------------------------------------------------------------------
 */

// DosenAccountIndex handler
func DosenAccountIndex(c echo.Context) error {
	id := c.Param("id")
	dosen, err := DB.Dosen(id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{
			"msg": "Not found",
		})
	}

	return c.JSON(http.StatusOK, dosen)
}

// DosenAccountStore handler
func DosenAccountStore(c echo.Context) error {
	id := c.FormValue("dosen_id")
	namaLengkap := c.FormValue("nama_lengkap")
	password := helpers.BcryptHashMake(c.FormValue("password"))

	insertError := DB.DosenCreate(id, namaLengkap, password)

	if insertError != nil {
		log.Println(insertError)
		return c.JSON(http.StatusNotModified, map[string]string{
			"msg": "Not modified",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"msg": "New dosen created.",
		"result": map[string]string{
			"id":           id,
			"nama_lengkap": namaLengkap,
		},
	})
}

// DosenAccountDelete handler
func DosenAccountDelete(c echo.Context) error {
	id := c.Param("id")
	_, deleteErr := DB.DosenDelete(id)

	if deleteErr != nil {
		log.Println(deleteErr)
		return c.NoContent(http.StatusNotModified)
	}

	return c.NoContent(http.StatusNoContent)
}

// DosenAccountLogin handler
func DosenAccountLogin(c echo.Context) error {
	id := c.FormValue("dosen_id")
	password := c.FormValue("password")

	dosen, err := DB.Dosen(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"msg":    "Not found",
			"status": false,
		})
	}

	passwordMatch := helpers.BcryptHashCompare(password, dosen.Password)
	if passwordMatch == false {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"msg":    "Login failed",
			"status": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   dosen,
	})
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
