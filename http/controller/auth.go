package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/data/types"
	"github.com/yzaimoglu/mitocho/utils/crypto"
	"net/http"
)

func (ctrl *Controller) PostAPILogin(c echo.Context) error {
	var req types.PostLoginRequest
	if err := c.Bind(&req); err != nil {
		return ctrl.Error(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return ctrl.Error(http.StatusBadRequest, err.Error())
	}

	if req.Password != req.PasswordRepeat {
		return ctrl.Error(http.StatusBadRequest, "passwords do not match")
	}

	user, err := ctrl.Svc.GetUserByEmail(req.Email)
	if err != nil {
		return ctrl.Error(http.StatusUnauthorized, "incorrect email or password")
	}

	match, err := crypto.CheckPassword(req.Password, user.Password)
	if err != nil {
		return ctrl.Error(http.StatusInternalServerError, err.Error())
	}

	if !match {
		return ctrl.Error(http.StatusUnauthorized, "incorrect email or password")
	}

	// TODO: Create PasetoV4 Token
	return nil
}

func (ctrl *Controller) PostAPILoginTOTP(c echo.Context) error {
	return nil
}
