package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New(health)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// TODO: usar um middleware do Echo para o healthcheck

	e.Logger.Fatal(e.Start(":8080"))

	// Consumer()
	// Producer()
}
