package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Message: "hello world!",
		})
	})

	sysChan := make(chan os.Signal, 2)
	signal.Notify(sysChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-sysChan
	gracefulCtx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()
	if err := e.Shutdown(gracefulCtx); err != nil {
		panic(err)
	}
}
