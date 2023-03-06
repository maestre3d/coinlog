package di

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog"
	"github.com/maestre3d/coinlog/messaging"
	"github.com/maestre3d/coinlog/messaging/kafka"
	"github.com/maestre3d/coinlog/storage/sql"
	clhttp "github.com/maestre3d/coinlog/transport/http"
)

type coinlogHTTPConfig struct {
	Application   coinlog.Config
	Server        clhttp.Config
	Database      sql.Config
	MessageBroker kafka.Config
}

type CoinlogHTTP struct {
	Config coinlogHTTPConfig
	Echo   *echo.Echo
	Bus    *messaging.Bus
}

func (h CoinlogHTTP) Start() {
	sysChan := make(chan os.Signal, 2)
	signal.Notify(sysChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := h.Echo.Start(h.Config.Server.ListenAddress); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	go h.Bus.Start()
	<-sysChan
	h.gracefulShutdown()
}

func (h CoinlogHTTP) gracefulShutdown() {
	gracefulCtx, cancel := context.WithTimeout(context.Background(), h.Config.Server.GracefulShutdownThreshold)
	defer cancel()
	if err := h.Echo.Shutdown(gracefulCtx); err != nil {
		panic(err)
	}
	h.Bus.Shutdown()
}
