package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestWeather_200(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/weather/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("140010")

	err := Weather(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWeather_404_WithUnknownID(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/weather/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("999999")

	err := Weather(c)
	if err != nil {
		t.Fatal(err)
	}
}
