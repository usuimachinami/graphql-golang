package main

import (
	"app/common"
	"app/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", handler.Hello())
	e.POST("/login", handler.Login())
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(common.SECRET_KEY)))
	r.POST("", handler.Restricted())

	e.Start(":5000")
}
