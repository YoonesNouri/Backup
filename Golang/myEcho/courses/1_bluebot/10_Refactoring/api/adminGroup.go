package api

import (
    "github.com/labstack/echo/v4"
    "api/handlers"
)

func AdminGroup(g *echo.Group) {
    g.GET("/main", handlers.MainAdmin)
}