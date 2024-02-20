package view

import "github.com/labstack/echo/v4"

func (view *View) GetHealthViewRoutes(group *echo.Group) {
	group.GET("health", view.Controller.GetViewHealth)
	group.GET("health/request", view.Controller.GetViewHealthRequest)
}
