package controller

import "github.com/yzaimoglu/mitocho/mitocho"

type Controller struct {
	Mitocho *mitocho.Mitocho
}

func NewController(mitocho *mitocho.Mitocho) *Controller {
	return &Controller{
		Mitocho: mitocho,
	}
}
