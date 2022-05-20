package main

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"timey/api"
	_ "timey/api"
	"timey/context"
	"timey/db"
	"timey/service"
)

func main() {
	e := echo.New()
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	context.Bind[echo.Echo]("echo", e)

	db.Init()
	service.InitCustomerRepo()
	service.InitSOWRepo()
	api.InitCustomers()
	api.InitSOWs()

	e.Logger.Fatal(e.Start(":8080"))
}
