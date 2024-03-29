package webserver

import (
	"context"
	"errors"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4/v2"
	"net/http"
	"os"
	"os/signal"
	"po/configs"
	"po/internal/webserver/middlewares"
	"po/routes"
	"time"
)

type Webserver struct{}

func New() *Webserver {
	return &Webserver{}
}

func (a *Webserver) Serve(ctx context.Context) error {
	c, err := configs.NewApp()

	if err != nil {
		return err
	}

	e := echo.New()
	e.Use(middlewares.Context)
	e.Use(apmechov4.Middleware())
	e.HideBanner = true
	e.HidePort = true

	routes.Register(e)

	color.Green("⇨ http server started on http://127.0.0.1:%v\n", c.Port)

	if err := e.Start(":" + c.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		e.Logger.Fatal("shutting down the server")
	}

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return e.Shutdown(cc)
}
