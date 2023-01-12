package http

import (
	"net/http"

	"github.com/maestre3d/coinlog/transport"

	"github.com/labstack/echo/v4"
)

type HealthcheckController struct{}

var _ Controller = HealthcheckController{}

func NewHealthcheckController() HealthcheckController {
	return HealthcheckController{}
}

func (l HealthcheckController) MapRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, transport.BasicResponse{
			Message: http.StatusText(http.StatusOK),
		})
	})
}

func (l HealthcheckController) MapVersionedRoutes(_ *echo.Group) {}
