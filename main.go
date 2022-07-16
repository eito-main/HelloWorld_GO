package main

import "github.com/labstack/echo/v4"

func main() {

	e := echo.New()

	//"/"を変えることでアクセス先のURLを変更できる

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))

}