// https://www.youtube.com/watch?v=WJulAmJhPkQ&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY&index=4
// https://www.youtube.com/redirect?event=video_description&redir_token=QUFFLUhqbEZTUHdHaGlYcGIyc1p3THg3NlJZQ09keHFuZ3xBQ3Jtc0trV2gybTdGemVuRmYzQnIxQnBWelBFaHBWWGJiSU56T2RIXzNLX0c1U2pxSDUyM0NNQzZkVFdyR2FLXzFPNlBEZnZMS3MtQWNxMnZKOHFMNzFhWGxNdWU2WUg5cEczYW53WGx6ZnNBNl9ucVFuSGlOWQ&q=https%3A%2F%2Fgithub.com%2Fverybluebot%2Fecho-server-tutorial%2Ftree%2Fpart_4_grouping_middlewares&v=WJulAmJhPkQ

package main


import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
 )

type Cat struct {
    Name        string    `json:"name"`
    Type        string    `json:"type"`
}

type Dog struct {
    Name        string    `json:"name"`
    Type        string    `json:"type"`
}

type Hamster struct {
    Name        string    `json:"name"`
    Type        string    `json:"type"`
}

func yallo(c echo.Context) error {
    return c.String(http.StatusOK, "yallo from the web side!")
}

func getCats(c echo.Context) error {
    catName := c.QueryParam("name")
    catType := c.QueryParam("type")

    dataType := c.Param("data")

    if dataType == "string" {
        return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand his type is: %s\n", catName, catType))
    }

    if dataType == "json" {
        return c.JSON(http.StatusOK, map[string]string{
            "name": catName,
            "type": catType,
        })
    }

    return c.JSON(http.StatusBadRequest, map[string]string{
        "error": "you need to lets us know if you want json or string data",
    })
}

func addCat(c echo.Context) error {
    cat := Cat{}

    defer c.Request().Body.Close()

    b, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        log.Printf("Failed reading the request body for addCats: %s\n", err)
        return c.String(http.StatusInternalServerError, "")
    }

    err = json.Unmarshal(b, &cat)
    if err != nil {
        log.Printf("Failed unmarshaling in addCats: %s\n", err)
        return c.String(http.StatusInternalServerError, "")
    }

    log.Printf("this is your cat: %#v\n", cat)
    return c.String(http.StatusOK, "we got your cat!")
}

func addDog(c echo.Context) error {
    dog := Dog{}

    defer c.Request().Body.Close()

    err := json.NewDecoder(c.Request().Body).Decode(&dog)
    if err != nil {
        log.Printf("Failed processing addDog request: %s\n", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    log.Printf("this is your dog: %#v", dog)
    return c.String(http.StatusOK, "we got your dog!")
}

func addHamster(c echo.Context) error {
    hamster := Hamster{}

    err := c.Bind(&hamster)
    if err != nil {
        log.Printf("Failed processing addHamster request: %s\n", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    log.Printf("this is your hamster: %#v", hamster)
    return c.String(http.StatusOK, "we got your hamster!")
}

func mainAdmin(c echo.Context) error {
    return c.String(http.StatusOK, "horay you are on the secret amdin main page!")
}

func main() {
    fmt.Println("Welcome to the server")

    e := echo.New()

    g := e.Group("/admin")

    // this logs the server interaction
    g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
    }))

    g.GET("/main", mainAdmin)

    e.GET("/", yallo)
    e.GET("/cats/:data", getCats)

    e.POST("/cats", addCat)
    e.POST("/dogs", addDog)
    e.POST("/hamsters", addHamster)

    e.Start(":8000")
}