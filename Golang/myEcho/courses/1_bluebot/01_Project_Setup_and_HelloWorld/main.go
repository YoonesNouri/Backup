// https://www.youtube.com/watch?v=wgW22wo6-SY&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY&index=1

// package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	// Create an Echo instance
// 	e := echo.New()
// 	// Define a route
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello web page! from C:/Users/yoone/myProjects/Golang/myEcho/yn-demo-echo/my_echo_course/01")
// 	})
// 	// Start the server
// 	e.Start(":8080")
// }

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create an Echo instance
	e := echo.New()
	// Define a route
	e.GET("/", func(c echo.Context) error {
		// Disable caching in the response headers
		c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("Expires", "0")

		return c.String(http.StatusOK, "Hello web page! from C:/Users/yoone/myProjects/Golang/myEcho/yn-demo-echo/my_echo_course/01")
	})
	// Start the server
	e.Start(":8000")
}
