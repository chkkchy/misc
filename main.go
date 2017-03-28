package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

const (
	hokkaido int = iota + 1
	aomori
	iwate
	miyagi
	akita
	yamagata
	fukushima
	ibaraki
	tochigi
	gunma
	saitama
	chiba
	toukyo
	kanagawa
	niigata
	toyama
	ishikawa
	fukui
	yamanashi
	nagano
	gifu
	sizuoka
	aichi
	mie
	shiga
	kyoto
	osaka
	hyogo
	nara
	wakayama
	tottori
	shimane
	okayama
	hiroshima
	yamaguchi
	tokushima
	kagawa
	ehime
	kochi
	fukuoka
	saga
	nagasaki
	kumamoto
	oita
	miyazaki
	kagoshima
	okinawa
)

// User is user
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

	e.GET("/weather", func(c echo.Context) error {
		fmt.Printf("%01d\n", 2)
		fmt.Println(okinawa)

		url := "http://weather.livedoor.com/forecast/webservice/json/v1?city=400040"
		req, _ := http.NewRequest("GET", url, nil)
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		type res struct {
			Title string `json:"title"`
			Link  string `json:"link"`
		}
		r := new(res)
		err = json.NewDecoder(resp.Body).Decode(r)
		return c.JSON(http.StatusOK, r)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
