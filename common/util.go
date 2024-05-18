package common

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func FormatValidationErrors(errors error) []string {
	return strings.Split(errors.Error(), "; ")
}

func Render(component templ.Component, context echo.Context, statusCode int) error {
	context.Response().WriteHeader(statusCode)
	return component.Render(context.Request().Context(), context.Response())
}

func RenderHandler(component templ.Component) echo.HandlerFunc {
	return func(context echo.Context) error {
		return Render(component, context, http.StatusOK)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
