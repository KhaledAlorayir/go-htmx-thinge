package middleware

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/constants"
	"github.com/labstack/echo/v4"
)

func RedirectIfAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(constants.AUTH_COOKIE_NAME)

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return next(c)
			}
			return err
		}

		_, err = common.ValidateJwt(cookie.Value)

		if err != nil {
			return err
		}

		return c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Protected(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(constants.AUTH_COOKIE_NAME)

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return c.Redirect(http.StatusMovedPermanently, constants.AUTH_PATH)
			}
			return err
		}

		data, err := common.ValidateJwt(cookie.Value)

		if err != nil {
			return err
		}

		userId, err := strconv.Atoi(data.Subject)

		if err != nil {
			return err
		}

		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), constants.CONTEXT_AUTH_DATA, common.AuthData{
			IsAuthenticated: true,
			Username:        data.Username,
			Email:           data.Email,
			Id:              userId,
		})))

		return next(c)
	}
}
