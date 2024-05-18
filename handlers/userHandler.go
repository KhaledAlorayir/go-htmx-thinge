package handlers

import (
	"net/http"

	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/dtos"
	"github.com/KhaledAlorayir/go-htmx-thinge/repository"
	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	repository repository.Repository
}

func NewUserHandler(repository repository.Repository) userHandler {
	return userHandler{
		repository: repository,
	}
}

func (h userHandler) CreateUserAction(c echo.Context) error {
	user, err := dtos.NewCreateUserRequest(c.FormValue("email"), c.FormValue("username"), c.FormValue("password"))

	if err != nil {
		return common.Render(views.ValidationErrors(common.FormatValidationErrors(err)), c, http.StatusBadRequest)
	}

	usernameExists, err := h.repository.CheckIfUsernameExists(user.Username)
	if err != nil {
		return err
	}
	emailExists, err := h.repository.CheckIfEmailExists(user.Email)
	if err != nil {
		return err
	}

	if usernameExists || emailExists {
		var messages []string
		if usernameExists {
			messages = append(messages, "username is already used!")
		}

		if emailExists {
			messages = append(messages, "email is already used!")
		}

		return common.Render(views.ValidationErrors(messages), c, http.StatusBadRequest)
	}

	user.Password, err = common.HashPassword(user.Password)

	if err != nil {
		return err
	}

	err = h.repository.CreateUser(user)

	if err != nil {
		return err
	}

	return common.Render(views.UserCreatedMessage(), c, http.StatusOK)
}
