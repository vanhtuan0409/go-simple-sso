package main

import (
	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web1/handler"
)

func main() {
	e := echo.New()
	h := handler.NewHandler()
	e.GET("/", h.Home)
	e.Start(":8081")
}
