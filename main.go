package main

import (
	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/handlers"
	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

type test struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	userHandler := handlers.UserHandler{}

	e.GET(common.CREATE_USER_ROUTE, handlers.RenderHandler(views.CreateUser()))
	e.POST(common.USER_PATH, userHandler.CreateUserAction)

	e.Logger.Fatal(e.Start(":3000"))
}
