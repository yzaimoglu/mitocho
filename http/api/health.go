package api

import (
	"github.com/labstack/echo/v4"
)

func (api *Api) GetHealthApiRoutes(group *echo.Group) {
	healthGroup := group.Group("health/")
	healthGroup.GET("", api.Controller.GetApiHealth)
}
