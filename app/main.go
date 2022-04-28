package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	// Echo instance
	e := echo.New()

	// Default routes
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// Start server with go routine
	go func() {
		address := fmt.Sprintf(":%d", 8080)
		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// Block until we receive our signal
	<-quit

	// Shutdown server
	if err := e.Shutdown(nil); err != nil {
		log.Fatal(err)
	}
}
