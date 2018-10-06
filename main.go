package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.Logger.Fatal(r.Start("localhost:1323"))
}
