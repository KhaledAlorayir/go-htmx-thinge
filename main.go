package main

import (
	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/constants"
	"github.com/KhaledAlorayir/go-htmx-thinge/handlers"
	"github.com/KhaledAlorayir/go-htmx-thinge/middleware"
	"github.com/KhaledAlorayir/go-htmx-thinge/repository"
	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	publicOnly := e.Group("", middleware.RedirectIfAuthenticated)
	protected := e.Group("", middleware.Protected)

	e.Debug = true
	db := initDatabase()
	repository := repository.NewRepository(db)

	userHandler := handlers.NewUserHandler(repository)
	authHandler := handlers.NewAuthHandler(repository)

	publicOnly.GET(constants.CREATE_USER_ROUTE, common.RenderHandler(views.CreateUserPage()))
	publicOnly.POST(constants.USER_PATH, userHandler.CreateUser)

	publicOnly.GET(constants.AUTH_PATH, common.RenderHandler(views.LoginPage()))
	publicOnly.POST(constants.AUTH_PATH, authHandler.Login)
	protected.POST(constants.LOGOUT_ROUTE, authHandler.Logout)

	protected.GET("/", common.RenderHandler(views.HomePage()))
	e.Logger.Fatal(e.Start(":3000"))
}
