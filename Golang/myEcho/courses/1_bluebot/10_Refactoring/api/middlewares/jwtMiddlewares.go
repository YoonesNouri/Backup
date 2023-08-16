package middlewares

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func SetJwtMiddlewares(g *echo.Group) {
    g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
        SigningMethod: "HS512",
        SigningKey: []byte("mySecret"),
    }))
}