package api

import "github.com/labstack/echo/v4"

func (api *Api) GetInstallApiRoutes(group *echo.Group) {
	installGroup := group.Group("install/")
	installGroup.POST("finish", api.Controller.PostAPIInstallFinish)
	installGroup.GET("finish", api.Controller.GetAPIInstallFinish)
}
