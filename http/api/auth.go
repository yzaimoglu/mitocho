package api

import "github.com/labstack/echo/v4"

func (api *Api) GetAuthApiRoutes(group *echo.Group) {
	authGroup := group.Group("auth/")
	authGroup.GET("login", nil)
	authGroup.GET("login/totp", nil)
	authGroup.GET("register", nil)
	authGroup.GET("register/totp", nil)
	authGroup.GET("", nil)
}
