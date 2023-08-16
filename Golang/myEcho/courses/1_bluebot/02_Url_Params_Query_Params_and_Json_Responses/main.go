// https://www.youtube.com/watch?v=wgW22wo6-SY&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY&index=2
// https://github.com/verybluebot/echo-server-tutorial/blob/part_2_params_params/src/main/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand his type is: %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to let us know if you want json or string data",
	})
}

func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()

	e.GET("/", yallo)
	e.GET("/cats/:data", getCats)

	e.Start(":8000")
}
