package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kabin-svvy/go-healthcheck-report/line/api/verify"
	"github.com/kabin-svvy/go-healthcheck-report/report"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(verify.LineJWT())

	e.POST("/healthcheck/report", report.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
