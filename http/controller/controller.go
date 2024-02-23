package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/data/services"
	"github.com/yzaimoglu/mitocho/mitocho"
)

type Controller struct {
	Mitocho *mitocho.Mitocho
	Svc     *services.Service
}

func NewController(mitocho *mitocho.Mitocho, svc *services.Service) *Controller {
	return &Controller{
		Mitocho: mitocho,
		Svc:     svc,
	}
}

func (ctrl *Controller) Error(code int, err string) error {
	return echo.NewHTTPError(code, err)
}
