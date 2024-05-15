package handlers

import (
	"net/http"

	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/dtos"
	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) CreateUserAction(c echo.Context) error {
	_, err := dtos.NewCreateUserRequest(c.FormValue("email"), c.FormValue("username"), c.FormValue("password"))

	if err != nil {
		return render(views.ValidationErrors(common.FormatValidationErrors(err)), c, http.StatusBadRequest)
	}

	return render(views.Message("hello!"), c, http.StatusOK)
}
