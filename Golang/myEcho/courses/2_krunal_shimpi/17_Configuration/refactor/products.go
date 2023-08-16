package refactor

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
)

// ProductsValidator is an echo validator for products
type ProductsValidator struct {
	validator *validator.Validate
}

// Validate validates the product request Body and returns an error if validation fails
func (p *ProductsValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "Product 1"}, {2: "Product 2"}, {3: "Product 3"}}

// getProducts returns the list of products
func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

// getProduct returns a specific product by ID
func getProduct(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if k == pID {
				return c.JSON(http.StatusOK, p)
			}
		}
	}
	return c.JSON(http.StatusNotFound, "product not found")
}

// createProduct creates a new product
func createProduct(c echo.Context) error {
	var product map[int]string
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	// Generating a unique ID for the new product
	newID := len(products) + 1
	product[newID] = product[1]
	products = append(products, product)
	return c.JSON(http.StatusCreated, product)
}

// updateProduct updates a product by ID
func updateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		for k := range p {
			if k == pID {
				product = p
				product[1] = "Updated Product" // Modify the product's name
				products[i] = product
				return c.JSON(http.StatusOK, product)
			}
		}
	}
	return c.JSON(http.StatusNotFound, "product not found")
}

// deleteProduct deletes a product by ID
func deleteProduct(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		for k := range p {
			if k == pID {
				products = append(products[:i], products[i+1:]...)
				return c.NoContent(http.StatusNoContent)
			}
		}
	}
	return c.JSON(http.StatusNotFound, "product not found")
}
