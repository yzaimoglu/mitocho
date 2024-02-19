package view

import (
	"github.com/yzaimoglu/mitocho/http/controller"
	"github.com/yzaimoglu/mitocho/mitocho"
)

type View struct {
	Mitocho    *mitocho.Mitocho
	Controller *controller.Controller
}

func NewView(mitocho *mitocho.Mitocho, controller *controller.Controller) *View {
	return &View{
		Mitocho:    mitocho,
		Controller: controller,
	}
}

func (view *View) Init() error {
	viewMain := view.Mitocho.Echo.Group("/")
	view.GetHealthViewRoutes(viewMain)
	return nil
}
