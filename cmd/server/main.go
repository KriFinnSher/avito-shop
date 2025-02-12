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
)

func main() {
	config.SetUpConfig()
	postgresDb := db.InitDb("postgres")
	db.MakeMigrations(true)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), "db", postgresDb)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})

	e.POST("/api/auth", handlers.AuthHandler)

	protectedGroup := e.Group("/api")
	protectedGroup.Use(mm.JwtMiddleware)

	protectedGroup.GET("/info", handlers.InfoHandler)
	protectedGroup.GET("/buy/:item", handlers.BuyItemHandler)
	protectedGroup.POST("/sendCoin", handlers.SendHandler)

	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
