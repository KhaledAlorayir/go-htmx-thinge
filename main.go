package main

import (
	"net/http"

	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

type test struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	e.GET("/user/create-user", func(c echo.Context) error {
		return views.CreateUser().Render(c.Request().Context(), c.Response())
	})

	e.POST("/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, test{Message: "yay"})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
