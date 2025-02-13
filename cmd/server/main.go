package main

import (
	"avito-shop/internal/config"
	"avito-shop/internal/db"
	"avito-shop/internal/handlers"
	mm "avito-shop/internal/middleware"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.SetUpConfig()
	postgresDB := db.InitDB("postgres")
	db.MakeMigrations(true)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), "db", postgresDB)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})

	e.POST("/api/auth", handlers.AuthHandler)

	protectedGroup := e.Group("/api")
	protectedGroup.Use(mm.JwtMiddleware)

	protectedGroup.GET("/info", handlers.InfoHandler)
	protectedGroup.GET("/buy/:item", handlers.BuyHandler)
	protectedGroup.POST("/sendCoin", handlers.SendHandler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start server")
		}
	}()

	<-stop
	slog.Info("received shutdown signal, starting shutdown...")

	// db.MakeMigrations(false)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		slog.Error("failed to gracefully shut down server", err.Error())
	}

	slog.Info("server gracefully stopped")
}
