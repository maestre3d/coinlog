package restapi

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/configuration"
)

type coinlogHTTPConfig struct {
	Application configuration.Application
	Server      configuration.ServerHTTP
	Database    configuration.DatabaseSQL
}

type CoinlogHTTP struct {
	Config coinlogHTTPConfig
	Echo   *echo.Echo
}

func (h CoinlogHTTP) Start() {
	sysChan := make(chan os.Signal, 2)
	signal.Notify(sysChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := h.Echo.Start(h.Config.Server.ListenAddress); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-sysChan
	h.gracefulShutdown()
}

func (h CoinlogHTTP) gracefulShutdown() {
	gracefulCtx, cancel := context.WithTimeout(context.TODO(), h.Config.Server.GracefulShutdownThreshold)
	defer cancel()
	if err := h.Echo.Shutdown(gracefulCtx); err != nil {
		panic(err)
	}
}
