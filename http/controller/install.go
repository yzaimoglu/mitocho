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

	if req.Password != req.PasswordRepeat {
		return ctrl.BadRequest(c, errors.New("passwords do not match"))
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

func (ctrl *Controller) GetAPIInstallFinish(c echo.Context) error {
	return ctrl.Success(c, types.JSON{
		"finished": ctrl.Svc.IsSetupFinished(),
	})
}
