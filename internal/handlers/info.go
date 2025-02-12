package handlers

import (
	"avito-shop/internal/models"
	"avito-shop/internal/repository/transaction"
	"avito-shop/internal/repository/user"
	"avito-shop/internal/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Coins     uint64
	Inventory map[string]uint64
	History   []models.Transaction
}

func InfoHandler(ctx echo.Context) error {
	db := ctx.Request().Context().Value("db").(*sqlx.DB)
	username, ok := ctx.Get("username").(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token or missing username")
	}

	reqCtx := ctx.Request().Context()

	userRepo := user.NewRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	transactionRepo := transaction.NewRepo(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)

	user_, exist := userUsecase.Exist(reqCtx, username)
	if !exist {
		return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
	}

	coins, err := userUsecase.GetBalance(reqCtx, user_.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get balance")
	}

	inventory, err := userUsecase.GetInventory(reqCtx, user_.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get inventory")
	}

	fmt.Println(user_.Name, "BIG BOY")
	history, err := transactionUsecase.GetHistory(reqCtx, user_.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get history")
	}

	response := Response{
		Coins:     coins,
		Inventory: inventory,
		History:   history,
	}

	return ctx.JSON(http.StatusOK, response)
}
