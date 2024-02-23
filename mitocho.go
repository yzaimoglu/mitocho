package main

import (
	"github.com/yzaimoglu/mitocho/config"
	"github.com/yzaimoglu/mitocho/data/services"
	"github.com/yzaimoglu/mitocho/http/api"
	"github.com/yzaimoglu/mitocho/http/controller"
	"github.com/yzaimoglu/mitocho/http/view"
	"github.com/yzaimoglu/mitocho/mitocho"
)

func main() {
	config.Load()
	db := config.NewDB()

	mtch := mitocho.NewMitocho(db)
	svc := services.NewService(db)
	ctrl := controller.NewController(mtch, svc)
	apiRouter := api.NewApi(mtch, ctrl)
	viewRouter := view.NewView(mtch, ctrl)

	err := apiRouter.Init()
	if err != nil {
		mtch.Echo.Logger.Fatalf("Api initialization failed: %v", err)
	}
	err = viewRouter.Init()
	if err != nil {
		mtch.Echo.Logger.Fatalf("View initialization failed: %v", err)
	}
	err = mtch.Start()
	if err != nil {
		mtch.Echo.Logger.Fatalf("Server start failed: %v", err)
	}
}
