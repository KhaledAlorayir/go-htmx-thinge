package main

import (
	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/handlers"
	"github.com/KhaledAlorayir/go-htmx-thinge/repository"
	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true
	db := initDatabase()
	repository := repository.NewRepository(db)

	userHandler := handlers.NewUserHandler(repository)
	authHandler := handlers.NewAuthHandler(repository)

	e.GET(common.CREATE_USER_ROUTE, common.RenderHandler(views.CreateUserPage()))
	e.POST(common.USER_PATH, userHandler.CreateUserAction)

	e.GET(common.AUTH_PATH, common.RenderHandler(views.LoginPage()))
	e.POST(common.AUTH_PATH, authHandler.Login)

	e.Logger.Fatal(e.Start(":3000"))
}
