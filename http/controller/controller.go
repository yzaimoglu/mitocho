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

type Response struct {
	Status    int    `json:"status"`
	RequestId string `json:"request_id"`
	Response  any    `json:"response,omitempty"`
}

func (ctrl *Controller) Response(c echo.Context, code int, data any) error {
	return c.JSON(code, Response{
		Status:    code,
		RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
		Response:  data,
	})
}

func (ctrl *Controller) Success(c echo.Context, data any) error {
	return ctrl.Response(c, http.StatusOK, data)
}

func (ctrl *Controller) BadRequest(c echo.Context, err error) error {
	return ctrl.Response(c, http.StatusBadRequest, types.JSON{
		"error": err.Error(),
	})
}

func (ctrl *Controller) InternalServerError(c echo.Context, err error) error {
	return ctrl.Response(c, http.StatusInternalServerError, types.JSON{
		"error": err.Error(),
	})
}

func (ctrl *Controller) ServiceUnavailable(c echo.Context, err error) error {
	return ctrl.Response(c, http.StatusServiceUnavailable, types.JSON{
		"error": err.Error(),
	})
}

func (ctrl *Controller) Unauthorized(c echo.Context, err error) error {
	return ctrl.Response(c, http.StatusUnauthorized, types.JSON{
		"error": err.Error(),
	})
}

func (ctrl *Controller) Forbidden(c echo.Context, err error) error {
	return ctrl.Response(c, http.StatusForbidden, types.JSON{
		"error": err.Error(),
	})
}

func (ctrl *Controller) NotFound(c echo.Context, err error) error {
	return ctrl.Response(c, http.StatusNotFound, types.JSON{
		"error": err.Error(),
	})
}
