package main

import (
	"fmt"

	"github.com/chkkchy/misc/api"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("vim-go")

	e := echo.New()

	e.GET("/weather/:id", api.Weather)

	e.Logger.Fatal(e.Start(":1323"))
}
