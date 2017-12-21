package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"graphql-golang/common"
	"graphql-golang/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", handler.Hello())
	e.POST("/login", handler.Login())
	r := e.Group("/query")
	r.Use(middleware.JWT([]byte(common.SECRET_KEY)))
	r.POST("", handler.Query())

	e.Start(":5000")
}
