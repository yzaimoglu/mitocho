package controller

import (
	"context"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/config"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

func (ctrl *Controller) View(c echo.Context, cmp templ.Component) error {
	ctrl.saveLocallyDev(cmp, c.Request().URL)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func (ctrl *Controller) saveLocallyDev(cmp templ.Component, url *url.URL) error {
	if config.Debug() {
		rootPath := "./templ/html/"

		dir := path.Join(rootPath)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			log.Fatalf("failed to create dir %q: %v", dir, err)
		}
		fileName := crypto.Sha256(url.Path) + ".html"

		filePath := path.Join(dir, fileName)
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}

		return cmp.Render(context.Background(), f)
	}
	return nil
}
