package api

import (
	"github.com/labstack/echo/v4"
)

func (api *Api) GetHealthApiRoutes(group *echo.Group) {
	group.GET("health", api.Controller.GetApiHealth)
}
