package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LivenessHTTP struct {
}

var _ HTTP = LivenessHTTP{}

func NewLivenessHTTP() LivenessHTTP {
	return LivenessHTTP{}
}

func (l LivenessHTTP) MapEndpoints(e *echo.Echo) {
	e.GET("/liveness", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}

func (l LivenessHTTP) MapVersionedEndpoints(_ *echo.Group) {
}
