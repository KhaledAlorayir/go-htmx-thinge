package handlers

import (
	"fmt"
	"net/http"

	"github.com/KhaledAlorayir/go-htmx-thinge/common"
	"github.com/KhaledAlorayir/go-htmx-thinge/repository"
	"github.com/KhaledAlorayir/go-htmx-thinge/views"
	"github.com/labstack/echo/v4"
)

type exerciseHandler struct {
	repository repository.Repository
}

func NewExerciseHandler(repository repository.Repository) exerciseHandler {
	return exerciseHandler{
		repository: repository,
	}
}

func (h exerciseHandler) HomePage(c echo.Context) error {
	data, error := h.repository.GetMuscleGroupsWithExercises()

	if error != nil {
		return error
	}
	fmt.Println(data)
	return common.Render(views.HomePage(data), c, http.StatusOK)
}
