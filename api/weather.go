package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

const urlFmt = "http://weather.livedoor.com/forecast/webservice/json/v1?city=%s" // e.g. 140010

type res struct {
	PinpointLocations []struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"pinpointLocations"`
	Link      string `json:"link"`
	Forecasts []struct {
		DateLabel   string `json:"dateLabel"`
		Telop       string `json:"telop"`
		Date        string `json:"date"`
		Temperature struct {
			Min interface{} `json:"min"`
			Max struct {
				Celsius    string `json:"celsius"`
				Fahrenheit string `json:"fahrenheit"`
			} `json:"max"`
		} `json:"temperature"`
		Image struct {
			Width  int    `json:"width"`
			URL    string `json:"url"`
			Title  string `json:"title"`
			Height int    `json:"height"`
		} `json:"image"`
	} `json:"forecasts"`
	Location struct {
		City       string `json:"city"`
		Area       string `json:"area"`
		Prefecture string `json:"prefecture"`
	} `json:"location"`
	PublicTime string `json:"publicTime"`
	Copyright  struct {
		Provider []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"provider"`
		Link  string `json:"link"`
		Title string `json:"title"`
		Image struct {
			Width  int    `json:"width"`
			Link   string `json:"link"`
			URL    string `json:"url"`
			Title  string `json:"title"`
			Height int    `json:"height"`
		} `json:"image"`
	} `json:"copyright"`
	Title       string `json:"title"`
	Description struct {
		Text       string `json:"text"`
		PublicTime string `json:"publicTime"`
	} `json:"description"`
}

// Weather gets weather info
func Weather(c echo.Context) error {
	id := c.Param("id")

	url := fmt.Sprintf(urlFmt, id)
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := new(res)
	json.NewDecoder(resp.Body).Decode(r)
	if r.Title == "" {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, r)
}
