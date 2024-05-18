package common

import (
	"net/http"
	"strings"
	"time"

	"github.com/KhaledAlorayir/go-htmx-thinge/repository"
	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
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

func GenerateJwt(user repository.User) (JWT, error) {
	expiresAt := time.Now().Add(time.Hour * 2)

	jwt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": expiresAt.Unix(),
	}).SignedString([]byte("secret"))

	return JWT{Jwt: jwt, ExpiresAt: expiresAt}, err
}
