package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ProductsValidator echo validator for products
type ProductsValidator struct {
	validator *validator.Validate
}

// validate validates the product request Body and returns
func (p *ProductsValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	v := validator.New()
	products := []map[int]string{{1: "Product 1"}, {2: "Product 2"}, {3: "Product 3"}}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range products {
			for k := range p {
				if k == pID {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
			//* the string field is required and must be at least 4 characters long
		}
		var reqBody body
		e.Validator = &ProductsValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := c.Validate(reqBody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusCreated, product)
	})

	e.PUT("/products/:id", func(c echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range products {
			for k := range p {
				if k == pID {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"` //* یعنی تو ورودی دیتای جیسون بنویس پروداکت نیم و برای تاییدش هم یه استرینگ با حداقل 4 کاراکتر باید باشه
		}
		var reqBody body
		e.Validator = &ProductsValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		product[pID] = reqBody.Name
		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
