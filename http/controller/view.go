package controller

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func (controller *Controller) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
