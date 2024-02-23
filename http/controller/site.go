package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/data/types"
	"net/http"
)

func (ctrl *Controller) GetAPISitePublicKey(c echo.Context) error {
	var req types.GetSitePublicKeyRequest
	if c.Bind(&req) != nil {
		return ctrl.Error(http.StatusBadRequest, "invalid request")
	}

	if c.Validate(&req) != nil {
		return ctrl.Error(http.StatusBadRequest, "invalid request")
	}

	site, err := ctrl.Svc.GetSiteByAccessToken(types.AccessToken(req.AccessToken))
	if err != nil {
		return ctrl.Error(http.StatusUnauthorized, "invalid access token")
	}

	return c.JSON(http.StatusOK, types.JSON{
		"public_key": site.PublicKey,
	})
}
