package api

import "github.com/labstack/echo/v4"

func (api *Api) GetSiteApiRoutes(group *echo.Group) {
	siteGroup := group.Group("site/")
	siteGroup.POST("public_key", api.Controller.PostAPISitePublicKey)
}
