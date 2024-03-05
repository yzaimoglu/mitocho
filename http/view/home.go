package view

import "github.com/labstack/echo/v4"

func (view *View) GetHomeViewRoutes(group *echo.Group) {
	group.GET("", view.Controller.GetViewHome)
}
