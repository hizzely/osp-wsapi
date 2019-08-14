package routes

import (
	Handler "github.com/hizzely/osp-wsapi/handlers"
	"github.com/labstack/echo/v4"
)

// Register the defined routes
func Register(router *echo.Echo) {
	// -- /
	router.GET("/", Handler.HomeIndex)

	// -- /api
	apiRouter := router.Group("/api")

	// -- /api/dosen
	apiDosen := apiRouter.Group("/dosen")

	// -- /api/dosen/account
	apiDosenAccount := apiDosen.Group("/account")
	apiDosenAccount.GET("/:id", Handler.DosenAccountIndex)
	apiDosenAccount.POST("/", Handler.DosenAccountStore)
	apiDosenAccount.DELETE("/:id", Handler.DosenAccountDelete)
	apiDosenAccount.POST("/login", Handler.DosenAccountLogin)

	// -- /api/dosen/lecture/subject
	apiDosenLectureSubject := apiDosen.Group("/lecture/subject")
	apiDosenLectureSubject.GET("/:id", Handler.DosenLectureSubjectDetail)
	apiDosenLectureSubject.POST("/", Handler.DosenLectureSubjectStore)

	// -- /api/dosen/presence/session
	apiDosenPresenceSession := apiDosen.Group("/presence/session")
	apiDosenPresenceSession.GET("/:id", Handler.DosenPresenceSessionDetail)
	apiDosenPresenceSession.POST("/create_classroom", Handler.DosenPresenceSessionStore)
	apiDosenPresenceSession.PUT("/:id", Handler.DosenPresenceSessionUpdate)
	apiDosenPresenceSession.DELETE("/:id", Handler.DosenPresenceSessionDelete)
	apiDosenPresenceSession.PUT("/:id/refresh_code", Handler.DosenPresenceSessionRefreshCode)

	// -- /api/classroom
	apiRouter.GET("/classroom", Handler.ClassroomIndex)
	apiRouter.POST("/classroom", Handler.ClassroomStore)

	// -- /api/student
	apiStudent := apiRouter.Group("/student")

	// -- /api/student/account
	apiStudentAccount := apiStudent.Group("/account")
	apiStudent.POST("/account", Handler.StudentAccountStore)
	apiStudentAccount.GET("/:npm", Handler.StudentAccountDetail)
	apiStudentAccount.POST("/", Handler.StudentAccountStore)
	apiStudentAccount.DELETE("/:npm", Handler.StudentAccountDelete)
	apiStudentAccount.POST("/login", Handler.StudentAccountLogin)

	// -- /api/student/classroom/courses
	apiStudent.GET("/classroom/courses", Handler.StudentClassroomCoursesIndex)

	// -- /api/student/presence
	apiStudentPresence := apiStudent.Group("/presence")
	apiStudentPresence.GET("/:npm", Handler.StudentPresenceDetail)
	apiStudentPresence.GET("/history", Handler.StudentPresenceHistory)
	apiStudentPresence.POST("/code", Handler.StudentPresenceByCode)
	apiStudentPresence.POST("/rfid", Handler.StudentPresenceByRfid)
}
