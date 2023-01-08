package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/configuration"
)

// MuxHTTP sets up a routing of a set of HTTP controllers.
type MuxHTTP struct {
	majorVersion uint16
	apiPrefixFmt string
	controllers  []HTTP
}

func NewMux(cfg configuration.Application, httpCfg configuration.ServerHTTP) *MuxHTTP {
	return &MuxHTTP{
		majorVersion: cfg.MajorVersion,
		apiPrefixFmt: httpCfg.RootPathFormat,
	}
}

func (h *MuxHTTP) Add(ctrls ...HTTP) {
	h.controllers = append(h.controllers, ctrls...)
}

func (h *MuxHTTP) RegisterRoutes(e *echo.Echo) {
	g := e.Group(fmt.Sprintf(h.apiPrefixFmt, h.majorVersion))
	for _, ctrl := range h.controllers {
		ctrl.MapEndpoints(e)
		ctrl.MapVersionedEndpoints(g)
	}
}
