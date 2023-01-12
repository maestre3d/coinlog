package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog"
)

type Controller interface {
	MapRoutes(e *echo.Echo)
	MapVersionedRoutes(g *echo.Group)
}

type ControllerMapper struct {
	majorVersion uint16
	apiPrefixFmt string
	controllers  []Controller
}

func NewControllerMapper(cfg coinlog.Config, httpCfg Config) *ControllerMapper {
	return &ControllerMapper{
		majorVersion: cfg.MajorVersion,
		apiPrefixFmt: httpCfg.RootPathFormat,
	}
}

func (h *ControllerMapper) Add(cc ...Controller) {
	h.controllers = append(h.controllers, cc...)
}

func (h *ControllerMapper) RegisterRoutes(e *echo.Echo) {
	g := e.Group(fmt.Sprintf(h.apiPrefixFmt, h.majorVersion))
	for _, ctrl := range h.controllers {
		ctrl.MapRoutes(e)
		ctrl.MapVersionedRoutes(g)
	}
}
