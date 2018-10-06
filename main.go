package main

import (
	"log"

	"github.com/h00s/go-web-app/config"
	"github.com/h00s/go-web-app/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	c, err := config.Load("config/configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.New(c.Log.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	r := echo.New()
	r.Use(middleware.Static("public"))
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	l.Info("Starting server...")
	r.Logger.Fatal(r.Start(c.Server.IPAddress + ":" + c.Server.Port))
}
