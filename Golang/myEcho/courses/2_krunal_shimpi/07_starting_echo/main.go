package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error { //* receives a Context interface
		return c.String(http.StatusOK, "Hello, World!") //* returns a string of Context interface
	})
	e.Logger.Print("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
