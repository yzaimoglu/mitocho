package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/data/types"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

func (ctrl *Controller) PostAPISitePublicKey(c echo.Context) error {
	var req types.PostSitePublicKeyRequest
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

	pasetoGen := crypto.FromPrivateKey(string(site.PrivateKey))

	return ctrl.Success(c, types.JSON{
		"public_key": pasetoGen.ExportPublicKey(),
	})
}
