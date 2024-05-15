package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(component templ.Component, context echo.Context, statusCode int) error {
	context.Response().WriteHeader(statusCode)
	return component.Render(context.Request().Context(), context.Response())
}

func RenderHandler(component templ.Component) echo.HandlerFunc {
	return func(context echo.Context) error {
		return render(component, context, http.StatusOK)
	}
}
