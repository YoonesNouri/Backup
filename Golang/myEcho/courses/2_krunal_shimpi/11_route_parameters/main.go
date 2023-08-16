package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []string{"Product 1", "Product 2", "Product 3"})
	})

	e.GET("/products/:vendor", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.QueryParam("olderThan"))
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
