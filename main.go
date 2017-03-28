package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("vim-go")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		u := new(User)
		u.Name = "yukichi"
		u.Age = 13
		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
