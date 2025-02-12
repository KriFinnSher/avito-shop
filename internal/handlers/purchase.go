package handlers

import (
	"avito-shop/internal/repository/merch"
	"avito-shop/internal/repository/transaction"
	"avito-shop/internal/repository/user"
	"avito-shop/internal/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

func BuyItemHandler(ctx echo.Context) error {
	db := ctx.Request().Context().Value("db").(*sqlx.DB)
	username, ok := ctx.Get("username").(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token or missing username")
	}

	itemName := ctx.Param("item")

	reqCtx := ctx.Request().Context()

	userRepo := user.NewRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	itemRepo := merch.NewRepo(db)
	itemUsecase := usecase.NewMerchUsecase(itemRepo)

	transactionRepo := transaction.NewRepo(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)

	sender, exist := userUsecase.Exist(reqCtx, username)
	if !exist {
		return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
	}

	item, exist := itemUsecase.Exist(reqCtx, itemName)
	if !exist {
		return echo.NewHTTPError(http.StatusBadRequest, "item not found")
	}

	if sender.Balance < item.Cost {
		return echo.NewHTTPError(http.StatusBadRequest, "insufficient balance")
	}

	err := transactionUsecase.Purchase(reqCtx, sender, item)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to make transaction"})
	}

	err = userUsecase.UpdateBalance(reqCtx, sender.Id, -item.Cost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update user's balance")
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "item purchased successfully"})
}
