package api

import (
	"github.com/yzaimoglu/mitocho/http/controller"
	"github.com/yzaimoglu/mitocho/mitocho"
)

type Api struct {
	Mitocho    *mitocho.Mitocho
	Controller *controller.Controller
}

func NewApi(mitocho *mitocho.Mitocho, controller *controller.Controller) *Api {
	return &Api{
		Mitocho:    mitocho,
		Controller: controller,
	}
}

func (api *Api) Init() error {
	apiV1 := api.Mitocho.Echo.Group("/api/v1/")
	api.GetAuthApiRoutes(apiV1)
	api.GetSiteApiRoutes(apiV1)
	api.GetInstallApiRoutes(apiV1)
	return nil
}
