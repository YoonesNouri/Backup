package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Yallo(c echo.Context) error {
    return c.String(http.StatusOK, "yallo from the web side!")
}