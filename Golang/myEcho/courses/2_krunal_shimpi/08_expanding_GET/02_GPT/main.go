package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Map to store product information
	products := map[string]string{
		"1": "Mobile",
		"2": "TV",
		"3": "Laptop",
	}

	// Endpoint to get product by ID
	e.GET("/products/:id", func(c echo.Context) error {
		productID := c.Param("id")
		productName, found := products[productID]
		if !found {
			return echo.NewHTTPError(http.StatusNotFound, "Product not found")
		}
		return c.String(http.StatusOK, "This is a "+productName)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
