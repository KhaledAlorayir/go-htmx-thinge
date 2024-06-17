package common

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/KhaledAlorayir/go-htmx-thinge/constants"
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

func Redirect(path string, context echo.Context) error {
	context.Response().Header().Set("HX-Redirect", "/")
	return context.Redirect(200, "/")
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
	claims := &jwtData{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(user.Id),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	jwt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))

	return JWT{Jwt: jwt, ExpiresAt: expiresAt}, err
}

func ValidateJwt(token string) (jwtData, error) {
	claims := &jwtData{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	return *claims, err
}

func GetAuthData(context context.Context) AuthData {
	data := context.Value(constants.CONTEXT_AUTH_DATA)

	if data == nil {
		return AuthData{IsAuthenticated: false}
	}
	return data.(AuthData)
}

func ToJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
