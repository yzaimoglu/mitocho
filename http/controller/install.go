package controller

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/data/types"
)

func (ctrl *Controller) PostAPIInstallFinish(c echo.Context) error {
	var req types.PostInstallFinishRequest
	if c.Bind(&req) != nil {
		return ctrl.BadRequest(c, errors.New("could not bind request"))
	}

	if c.Validate(&req) != nil {
		return ctrl.BadRequest(c, errors.New("invalid request"))
	}

	if ctrl.Svc.IsSetupFinished() {
		return ctrl.Forbidden(c, errors.New("setup already finished"))
	}

	err := ctrl.Svc.FinishSetup(req.Name, req.Domain, req.Email, req.Username, req.Password)
	if err != nil {
		return ctrl.InternalServerError(c, errors.New("setup could not be finished"))
	}

	return ctrl.Success(c, nil)
}
