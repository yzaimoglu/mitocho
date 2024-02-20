package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/templ/health"
	"net/http"
)

func (ctrl *Controller) GetApiHealth(c echo.Context) error {
	return c.JSON(200, map[string]any{
		"status": "healthy",
	})
}

func (ctrl *Controller) GetViewHealth(c echo.Context) error {
	return ctrl.View(c,
		health.ShowIndex("test", "wow", health.Show("hey")))
}

func (ctrl *Controller) GetViewHealthRequest(c echo.Context) error {
	req := c.Request()
	format := `
		<code>
		  Protocol: %s<br>
		  Host: %s<br>
		  Remote Address: %s<br>
		  Method: %s<br>
		  Path: %s<br>
		</code>
	  `
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}
