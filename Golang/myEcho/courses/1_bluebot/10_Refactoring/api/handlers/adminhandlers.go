package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func MainAdmin(c echo.Context) error {
    return c.String(http.StatusOK, "horay you are on the secret amdin main page!")
}