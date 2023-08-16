package refactor

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.Readenv(&cfg)
	fmt.Printf("%+v", cfg )
	if err!= nil {
        e.Logger.Fatal("Unable to load configuration")
	}
}

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
fmt.Println("inside middleware")
c.Request().URL.Path="/yoones"
fmt.Printf("%+v", c.Request())
		return next(c)
    }

// * Start starts the application
func Start() {
	e.Use(RemoveTrailingSlash())
	//? e.Pre() // Pre middleware initialization code here to avoid race conditions when running multiple instances of the same application
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct, middleware.BodyLimit("1K"))
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
