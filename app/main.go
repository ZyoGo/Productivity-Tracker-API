package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/w33h/Productivity-Tracker-API/api"
	modules "github.com/w33h/Productivity-Tracker-API/app/module"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/w33h/Productivity-Tracker-API/config"
	"github.com/w33h/Productivity-Tracker-API/util"
)

func main() {
	// Init config
	cfg := config.GetConfig()

	dbCon := util.NewConnectionDB(cfg)
	defer dbCon.Close()

	controllers := modules.RegisterModules(dbCon)

	// Echo instance
	e := echo.New()

	// Default routes
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	api.RegistrationPath(e, controllers)

	// Start server with go routine
	go func() {
		address := fmt.Sprintf(":%d", cfg.App.Port)
		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown server
	quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)

	// Block until we receive our signal
	<-quit

	// timeout to wait for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
