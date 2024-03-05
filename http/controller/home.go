package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/templ/home"
)

func (ctrl *Controller) GetViewHome(c echo.Context) error {
	return ctrl.View(c,
		home.ShowIndex("Mitocho", "The powerhouse of authentication", home.Show()))
}
