package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("hello yoones")

	app := echo.New()

	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello yoones")
	})
	app.Logger.Fatal(app.Start(":8080"))

}

type data interface {
	test1()
	test2()
	test3()
	test4()
}
