package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/templ/home"
	"github.com/yzaimoglu/mitocho/templ/layout"
)

func (ctrl *Controller) GetViewHome(c echo.Context) error {
	return ctrl.View(c,
		layout.BaseView("Mitocho", "The powerhouse of authentication", home.Show()))
}
