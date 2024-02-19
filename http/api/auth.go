package api

import "github.com/labstack/echo/v4"

func (api *Api) GetAuthApiRoutes(group *echo.Group) {
	group.GET("login", nil)
	group.GET("login/totp", nil)
	group.GET("register", nil)
	group.GET("register/totp", nil)
	group.GET("", nil)
}
