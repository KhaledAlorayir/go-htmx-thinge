package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserRequest struct {
	Email    string
	Username string
	Password string
}

func NewCreateUserRequest(email string, username string, password string) (CreateUserRequest, error) {
	user := CreateUserRequest{Email: email, Username: username, Password: password}

	err := validation.ValidateStruct(
		&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Username, validation.Required, validation.Length(4, 40)),
		validation.Field(&user.Password, validation.Required, validation.Length(4, 500)),
	)

	if err != nil {
		return user, err
	}

	return user, nil
}
