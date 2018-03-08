package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "hello world!!!")
	})
	e.Start(":8081")
}
