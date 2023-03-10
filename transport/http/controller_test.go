package http_test

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog"
	"github.com/maestre3d/coinlog/transport/http"
	"github.com/stretchr/testify/assert"
)

type fakeController struct{}

var _ http.Controller = fakeController{}

func (f fakeController) MapRoutes(e *echo.Echo) {
	e.GET("/fake", func(c echo.Context) error {
		return nil
	})
}

func (f fakeController) MapVersionedRoutes(g *echo.Group) {
	g.DELETE("/fake", func(c echo.Context) error {
		return nil
	})
}

func TestNewControllerMapper(t *testing.T) {
	cfg := coinlog.NewConfig()
	cfg.MajorVersion = 12
	httpCfg := http.NewConfig()
	httpCfg.RootPathFormat = "/v%d"
	mapper := http.NewControllerMapper(cfg, httpCfg)
	mapper.Add(http.HealthcheckController{}, fakeController{})

	e := echo.New()
	mapper.RegisterRoutes(e)
	expMap := map[string]string{
		"/health":   "GET",
		"/fake":     "GET",
		"/v12/fake": "DELETE",
	}
	routes := e.Routes()
	foundCount := len(routes)
	for _, r := range routes {
		expRoute, ok := expMap[r.Path]
		if !ok {
			continue
		}
		foundCount--
		assert.Equal(t, expRoute, r.Method)
	}
	assert.Zero(t, foundCount)
}
