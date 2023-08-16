package main

// Import the required packages
import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Define the authentication function
func authenticate(username, password string, c echo.Context) (bool, error) {
    // Implement your authentication logic here
    if username == "jack" && password == "1234" {
        return true, nil
    }
    return false, nil
}

func main() {
    // Create a new Echo instance
    e := echo.New()

    // Use the BasicAuth middleware to enable browser-based authentication
    e.Use(middleware.BasicAuth(authenticate))

    // Define a protected route
    e.GET("/protected", func(c echo.Context) error {
        return c.String(http.StatusOK, "This is a protected route!")
    }, middleware.BasicAuth(authenticate))

    // Start the server
    e.Start(":8080")
}
