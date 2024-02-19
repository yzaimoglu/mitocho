package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func WriteCookie(c echo.Context, name string, value string, expires time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	c.SetCookie(cookie)
}

func ReadCookie(c echo.Context, name string) (*http.Cookie, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return nil, err
	}
	return cookie, err
}

func DeleteCookie(c echo.Context, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.MaxAge = -1
	c.SetCookie(cookie)
}
