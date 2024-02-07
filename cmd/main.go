package main

import (
	"logictest/database"
	"logictest/route"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

var (
	logger = middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} method=${method}, uri=${uri}, status=${status}, origin=${header:origin}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	})
)

func main() {

	e := echo.New()
	// e.GET("/", func(context echo.Context) error {
	// 	return context.String(http.StatusOK, "OK")
	// })
	e.Use(logger)
	// e.Use(cors)

	db, err := database.NewProgressSql("url")
	if err != nil {
		panic(err)
	}

	route.InitV1(e, db)
	e.Logger.Fatal(e.Start("0.0.0.0:4005"))
}
