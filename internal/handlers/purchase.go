package handlers

import (
	"avito-shop/internal/repository/postgre/merch"
	"avito-shop/internal/repository/postgre/transaction"
	"avito-shop/internal/repository/postgre/user"
	"avito-shop/internal/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

// BuyHandler creates a transaction and stores it in db if both item and user exist
// and balance of user is sufficient
func BuyHandler(ctx echo.Context) error {
	db := ctx.Request().Context().Value("db").(*sqlx.DB)
	username, ok := ctx.Get("username").(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{Message: "invalid token or missing username"})
	}

	itemName := ctx.Param("item")

	reqCtx := ctx.Request().Context()

	userRepo := user.NewPostgreRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	itemRepo := merch.NewPostgreRepo(db)
	itemUsecase := usecase.NewMerchUsecase(itemRepo)

	transactionRepo := transaction.NewPostgreRepo(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)

	sender, exist := userUsecase.Exist(reqCtx, username)
	if !exist {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{Message: "user not found"})
	}

	item, exist := itemUsecase.Exist(reqCtx, itemName)
	if !exist {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "item not found"})
	}

	if sender.Balance < item.Cost {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "insufficient balance"})
	}

	err := transactionUsecase.Purchase(reqCtx, &sender, &item)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to make transaction"})
	}

	err = userUsecase.UpdateBalance(reqCtx, sender.ID, -item.Cost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to update user's balance"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "item purchased successfully"})
}
