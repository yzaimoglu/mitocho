package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yzaimoglu/mitocho/templ/health"
	"net/http"
)

func (controller *Controller) GetApiHealth(c echo.Context) error {
	return c.JSON(200, map[string]any{
		"status": "healthy",
	})
}

func (controller *Controller) GetViewHealth(c echo.Context) error {
	return controller.View(c,
		health.ShowIndex("test", "wow", health.Show("hey")))
}

func (controller *Controller) GetApiHealthRequest(c echo.Context) error {
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
