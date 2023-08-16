package refactor

import ()

// ProductsValidator echo validator for products
type ProductsValidator struct {
	validator *validator.Validate
}

// validate validates the product request Body and returns
func (p *ProductsValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "Product 1"}, {2: "Product 2"}, {3: "Product 3"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
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
}

// POST method to create products
func createProduct(c echo.Context) error {
	var product map[int]string
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	products = append(products, product)
	return c.JSON(http.StatusCreated, product)
}

// PUT method to update products
func updateProduct(c echo.Context) error {
	var product map[int]string
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	for i, p := range products {
		if p["id"] == product["id"] {
			products[i] = product
		}
	}
	return c.JSON(http.StatusOK, product)
}

// DELETE method to delete products
func deleteProduct(c echo.Context) error {
	var product map[int]string
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	for i, p := range products {
		if p["id"] == product["id"] {
			products = append(products[:i], products[i+1:]...)
		}
	}
	return c.JSON(http.StatusOK, product)
}
