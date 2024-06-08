package handlers

import (
	"net/http"
	"time"

	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/constants"
	"github.com/KhaledAlorayir/go-htmx-thinge/dtos"
	"github.com/KhaledAlorayir/go-htmx-thinge/repository"
	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	repository repository.Repository
}

func NewAuthHandler(repository repository.Repository) authHandler {
	return authHandler{
		repository: repository,
	}
}

func (h authHandler) Login(c echo.Context) error {
	loginRequest, err := dtos.NewLoginRequest(c.FormValue("email"), c.FormValue("password"))

	if err != nil {
		return common.Render(views.ValidationErrors(common.FormatValidationErrors(err)), c, http.StatusBadRequest)
	}

	user, err := h.repository.GetUserByEmail(loginRequest.Email)

	if err != nil {
		return err
	}

	if user == nil {
		return common.Render(views.ValidationErrors([]string{"no user found with this email"}), c, http.StatusBadRequest)
	}

	authenticated := common.CheckPasswordHash(loginRequest.Password, user.Password)

	if !authenticated {
		return common.Render(views.ValidationErrors([]string{"invalid credentials"}), c, http.StatusBadRequest)
	}

	jwt, err := common.GenerateJwt(*user)

	if err != nil {
		return err
	}

	setAuthCookie(c, jwt)
	return common.Redirect("/", c)
}

func (h authHandler) Logout(c echo.Context) error {
	deleteAuthCookie(c)
	return common.Redirect("/auth", c)
}

func setAuthCookie(context echo.Context, jwt common.JWT) {
	cookie := new(http.Cookie)
	cookie.Name = constants.AUTH_COOKIE_NAME
	cookie.Value = jwt.Jwt
	cookie.Expires = jwt.ExpiresAt
	cookie.HttpOnly = true
	cookie.Secure = true
	context.SetCookie(cookie)
}

func deleteAuthCookie(context echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = constants.AUTH_COOKIE_NAME
	cookie.Expires = time.Unix(0, 0)
	cookie.Value = ""
	cookie.Path = "/"
	context.SetCookie(cookie)
}
