package refactor

import (
	"fmt"
	"os"

	"github.com/go_playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var v = validator.New()

// * Start starts the application
func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.Get("/products", getProducts)
	e.Get("/products/:id", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
