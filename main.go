package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/poccariswet/shorterql/handler"
)

func main() {
	e := echo.New()
	e.Validator = &handler.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.SetLevel(log.INFO)

	e.GET("/:id", handler.RedirectHandler)
	e.GET("urlshorter/status/:id", handler.UrlShorterStatusHandler)
	e.POST("urlshorter", handler.UrlShorterHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Info("shut down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
