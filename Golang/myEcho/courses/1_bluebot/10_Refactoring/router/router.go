//? https://www.youtube.com/watch?v=S2PfH0zA2Vw&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY&index=11
//? https://github.com/verybluebot/echo-server-tutorial/tree/part_10_refactoring

package router

import (
    "api/middlewares"
    "github.com/labstack/echo/v4"
    "api"
)

func New() *echo.Echo {
    e := echo.New()

    // create groups
    adminGroup := e.Group("/admin")
    cookieGroup := e.Group("/cookie")
    jwtGroup := e.Group("/jwt")

    // set all middlewares
    middlewares.SetMainMiddlewares(e)
    middlewares.SetAdminMiddlewares(adminGroup)
    middlewares.SetCookieMiddlewares(cookieGroup)
    middlewares.SetJwtMiddlewares(jwtGroup)

    // set main routes
    api.MainGroup(e)

    // set group routes
    api.AdminGroup(adminGroup)
    api.CookieGroup(cookieGroup)
    api.JwtGroup(jwtGroup)

    return e
}