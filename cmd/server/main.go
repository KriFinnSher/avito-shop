package main

import (
	"avito-shop/internal/config"
	"avito-shop/internal/db"
	//"avito-shop/internal/handlers"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
)

func main() {
	config.SetUpConfig()
	postgresDb := db.InitDb("postgres")
	db.MakeMigrations()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.POST("/api/auth", handlers.Auth)
	//e.GET("/api/info", handlers.GetInfo)
	//e.GET("/api/buy/:item", handlers.BuyItem)
	//e.POST("/api/sendCoin", handlers.SendCoins)
	e.GET("/api/test", SelectHandler(postgresDb))

	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}

}

type price uint64

type Item struct {
	Id   uint64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Cost price  `db:"cost" json:"cost"`
}

func SelectAll(db *sqlx.DB) []Item {
	var items []Item
	db.Select(&items, "SELECT * FROM merch")
	return items
}

func SelectHandler(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		items := SelectAll(db)
		fmt.Println(items)
		return nil
	}
}
