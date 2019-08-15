package handlers

import (
	"net/http"

	DB "github.com/hizzely/osp-wsapi/database"
	"github.com/labstack/echo/v4"
)

/*
 * ------------------------------------------------------------------
 * Classroom
 * ------------------------------------------------------------------
 */

// ClassroomIndex handler
func ClassroomIndex(c echo.Context) error {
	kelas, _ := DB.ClassroomAll()
	return c.JSON(http.StatusOK, kelas)
}

// ClassroomStore handler
func ClassroomStore(c echo.Context) error {
	namaKelas := c.FormValue("nama_kelas")
	err := DB.ClassroomCreate(namaKelas)
	status := true

	if err != nil {
		status = false
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"result": status,
		"nama_kelas": namaKelas,
	})
}
