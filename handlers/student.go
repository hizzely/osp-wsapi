package handlers

import (
	"net/http"
	"strconv"
	"log"

	"github.com/labstack/echo/v4"
	DB "github.com/hizzely/osp-wsapi/database"
	"github.com/hizzely/osp-wsapi/helpers"
)

/*
 * ------------------------------------------------------------------
 * Student / Account
 * ------------------------------------------------------------------
 */

// StudentAccountDetail handler
func StudentAccountDetail(c echo.Context) error {
	id := c.Param("npm")
	student, err := DB.Student(id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{
			"msg": "Not found",
		})
	}

	return c.JSON(http.StatusOK, student)
}

// StudentAccountStore handler
func StudentAccountStore(c echo.Context) error {
	id := c.FormValue("npm")
	namaLengkap := c.FormValue("nama_lengkap")
	password := helpers.BcryptHashMake(c.FormValue("password"))
	kelasID, _ := strconv.Atoi(c.FormValue("kelas_id"))
	status := "aktif"

	resultErr := DB.StudentCreate(id, namaLengkap, password, status, kelasID)
	
	if resultErr != nil {
		log.Println(resultErr)
		return c.JSON(http.StatusNotModified, map[string]string{
			"msg": "Not modified",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"msg": "New student created.",
		"result": map[string]interface{} {
			"id": id,
			"nama_lengkap": namaLengkap,
			"kelas_id": kelasID,
		},
	})
}

// StudentAccountDelete handler
func StudentAccountDelete(c echo.Context) error {
	id := c.Param("npm")
	_, deleteErr := DB.StudentDelete(id)

	if deleteErr != nil {
		log.Println(deleteErr)
		return c.NoContent(http.StatusNotModified)
	}

	return c.NoContent(http.StatusNoContent)
}

// StudentAccountLogin handler
func StudentAccountLogin(c echo.Context) error {
	id := c.FormValue("npm")
	password := c.FormValue("password")

	student, err := DB.Student(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"msg":    "Not found",
			"status": false,
		})
	}

	passwordMatch := helpers.BcryptHashCompare(password, student.Password)
	if passwordMatch == false {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"msg":    "Login failed",
			"status": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data": student,
	})
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
	studentID := c.FormValue("npm")
	matkulID, err := strconv.Atoi(c.FormValue("matkul_id"))

	if err != nil {
		matkulID = 0
	}

	result, err := DB.StudentPresenceHistory(studentID, matkulID)
	return c.JSON(http.StatusOK, result)
}

// StudentPresenceByCode handler
func StudentPresenceByCode(c echo.Context) error {
	id := c.FormValue("npm")
	sessionCode := c.FormValue("session_code")
	result := DB.StudentPresenceByCode(id, sessionCode)

	if result == 1 {
		return c.JSON(http.StatusCreated, map[string]string {
			"msg": "Presence success",
		})
	} else if result == 2 {
		return c.JSON(http.StatusOK, map[string]string {
			"msg": "Already present",
		})
	} else {
		return c.JSON(http.StatusNotFound, map[string]string {
			"msg": "Session data does not exist",
		})
	}
}

// StudentPresenceByRfid handler
func StudentPresenceByRfid(c echo.Context) error {
	rfidCode := c.FormValue("rfid_code")
	sessionCode := c.FormValue("session_code")
	result := DB.StudentPresenceByRfid(rfidCode, sessionCode)

	if result == 1 {
		return c.JSON(http.StatusCreated, map[string]string {
			"msg": "Presence success",
		})
	} else if result == 2 {
		return c.JSON(http.StatusOK, map[string]string {
			"msg": "Already present",
		})
	} else {
		return c.JSON(http.StatusNotFound, map[string]string {
			"msg": "Session data does not exist",
		})
	}
}
