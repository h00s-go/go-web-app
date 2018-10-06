package main

import (
	"github.com/h00s/go-web-app/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	c, err := config.Load("config/configuration.json")
	if err != nil {
		panic(err)
	}

	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.Logger.Fatal(r.Start(c.Server.IPAddress + ":" + c.Server.Port))
}
