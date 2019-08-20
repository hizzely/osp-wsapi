package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

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
	dosenID := c.Param("id")
	result, _ := DB.DosenLectureSubject(dosenID)

	return c.JSON(http.StatusOK, result)
}

// DosenLectureSubjectStore handler
func DosenLectureSubjectStore(c echo.Context) error {
	dosenID := c.Param("dosen_id")
	matkulID, _ := strconv.Atoi(c.Param("matkul_id"))
	kelasID, _ := strconv.Atoi(c.Param("kelas_id"))

	insertErr := DB.DosenLectureSubjectCreate(dosenID, matkulID, kelasID)

	if insertErr != nil {
		log.Println(insertErr)
		return c.JSON(http.StatusNotModified, map[string]string{
			"msg": "Not modified",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"msg": "New lecture subject created.",
	})
}

/*
 * ------------------------------------------------------------------
 * Dosen / Presence
 * ------------------------------------------------------------------
 */

// DosenPresenceSessionDetail handler
func DosenPresenceSessionDetail(c echo.Context) error {
	presenceID, _ := strconv.Atoi(c.Param("id"))
	details := DB.DosenPresenceSessionDetail(presenceID)

	return c.JSON(http.StatusOK, details)
}

// DosenPresenceSessionStore handler
func DosenPresenceSessionStore(c echo.Context) error {
	dosenID := c.FormValue("dosen_id")
	matkulID, _ := strconv.Atoi(c.FormValue("matkul_id"))
	kelasID, _ := strconv.Atoi(c.FormValue("kelas_id"))
	judul := c.FormValue("judul")
	deskripsi := c.FormValue("deskripsi")
	sessionCode, _ := helpers.GenerateRandomString(6)

	// If at some point the session code already used, retry.
	for i := 0; i < 3; i++ {
		resultErr := DB.DosenPresenceSessionCreateClassroom(matkulID, kelasID, dosenID, judul, deskripsi, sessionCode)
		if resultErr == nil {
			break
		}
		sessionCode, _ = helpers.GenerateRandomString(6)
	}

	kelas, _ := DB.Classroom(kelasID)
	matkul, _ := DB.Matkul(matkulID)
	dosen, _ := DB.Dosen(dosenID)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id_dosen":    dosenID,
		"nama_dosen":  dosen.NamaLengkap,
		"id_matkul":   matkulID,
		"nama_matkul": matkul.NamaMatkul,
		"id_kelas":    kelasID,
		"nama_kelas":  kelas.NamaKelas,
		"kode":        sessionCode,
		"start":       time.Now(),
		"judul":       judul,
		"deskripsi":   deskripsi,
	})
}

// DosenPresenceSessionUpdate handler
func DosenPresenceSessionUpdate(c echo.Context) error {
	sessionID, _ := strconv.Atoi(c.Param("id"))
	judul := c.FormValue("judul")
	deskripsi := c.FormValue("deskripsi")
	status := c.FormValue("status")
	result, err := DB.DosenPresenceSessionUpdate(sessionID, judul, deskripsi, status)

	if err != nil || result == 0 {
		log.Println(err)
		return c.JSON(http.StatusNotModified, "Not modified")
	}

	return c.JSON(http.StatusOK, "Success")
}

// DosenPresenceSessionDelete handler
func DosenPresenceSessionDelete(c echo.Context) error {
	sessionID, _ := strconv.Atoi(c.Param("id"))
	result, err := DB.DosenPresenceSessionDelete(sessionID)

	if err != nil || result == 0 {
		log.Println(err)
		return c.JSON(http.StatusNotModified, "Not modified")
	}

	return c.JSON(http.StatusOK, "Success")
}

// DosenPresenceSessionRefreshCode handler
func DosenPresenceSessionRefreshCode(c echo.Context) error {
	sessionID, _ := strconv.Atoi(c.Param("id"))
	randomCode, _ := helpers.GenerateRandomString(6)
	result, err := DB.DosenPresenceSessionRefreshCode(sessionID, randomCode)

	if err != nil || result == 0 {
		log.Println(err)
		return c.JSON(http.StatusNotModified, "Not modified")
	}

	return c.JSON(http.StatusOK, map[string]string {
		"status": "success",
		"kode": randomCode,
	})
}
