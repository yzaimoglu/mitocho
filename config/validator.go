package config

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MitochoValidator struct {
	validator *validator.Validate
}

func NewValidator() *MitochoValidator {
	return &MitochoValidator{
		validator: validator.New(),
	}
}

func (mv *MitochoValidator) Validate(i interface{}) error {
	if err := mv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
