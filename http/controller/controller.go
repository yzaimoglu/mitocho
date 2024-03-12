package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/data/services"
	"github.com/yzaimoglu/mitocho/data/types"
	"github.com/yzaimoglu/mitocho/mitocho"
)

type Controller struct {
	Mitocho *mitocho.Mitocho
	Svc     *services.Service
}

func NewController(mitocho *mitocho.Mitocho, svc *services.Service) *Controller {
	return &Controller{
		Mitocho: mitocho,
		Svc:     svc,
	}
}

func (ctrl *Controller) Error(code int, err string) error {
	return echo.NewHTTPError(code, err)
}

func (ctrl *Controller) Response(c echo.Context, code int, data any) error {
	return c.JSON(code, types.JSON{
		"status":     code,
		"request_id": c.Response().Header().Get(echo.HeaderXRequestID),
		"response":   data,
	})
}

func (ctrl *Controller) Success(c echo.Context, data any) error {
	return ctrl.Response(c, http.StatusOK, data)
}

func (ctrl *Controller) InternalServerError(c echo.Context, data any) error {
	return ctrl.Response(c, http.StatusInternalServerError, data)
}

func (ctrl *Controller) ServiceUnavailable(c echo.Context, data any) error {
	return ctrl.Response(c, http.StatusServiceUnavailable, data)
}

func (ctrl *Controller) Unauthorized(c echo.Context, data any) error {
	return ctrl.Response(c, http.StatusUnauthorized, data)
}

func (ctrl *Controller) Forbidden(c echo.Context, data any) error {
	return ctrl.Response(c, http.StatusForbidden, data)
}

func (ctrl *Controller) NotFound(c echo.Context, data any) error {
	return ctrl.Response(c, http.StatusForbidden, data)
}
