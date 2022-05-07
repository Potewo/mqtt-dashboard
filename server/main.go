package main

import (
	"fmt"
	"net/http"

	"github.com/Potewo/mqtt-dashboard/golib-db"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Printf("%#v\n", db.Atemperature)
	fmt.Println("success")
	temperatures := []db.Temperature{}
	db.DB.Find(&temperatures)
	fmt.Printf("\n%#v\n", temperatures)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
