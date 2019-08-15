package main

//go:generate sqlboiler -c=config/db.json --wipe mysql

import (
	"context"
	"database/sql"

	"github.com/hizzely/osp-wsapi/database"
	Routes "github.com/hizzely/osp-wsapi/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "DSN_STRING")
	if err != nil {
		panic(err)
	}
	database.Db = db
	database.Ctx = context.Background()

	// Echo instance
	echo := echo.New()
	echo.HideBanner = true

	// Middleware
	echo.Use(middleware.CORS())
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	// Routes
	Routes.Register(echo)

	// Start server
	echo.Logger.Fatal(echo.Start(":8200"))
}
