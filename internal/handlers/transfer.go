package handlers

import (
	"avito-shop/internal/repository/postgre/transaction"
	"avito-shop/internal/repository/postgre/user"
	"avito-shop/internal/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SendCoinRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

func SendHandler(ctx echo.Context) error {
	db := ctx.Request().Context().Value("db").(*sqlx.DB)
	username, ok := ctx.Get("username").(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{Message: "invalid token or missing username"})
	}

	var req SendCoinRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "Invalid request data"})
	}

	reqCtx := ctx.Request().Context()

	userRepo := user.NewPostgreRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	transactionRepo := transaction.NewPostgreRepo(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)

	sender, exist := userUsecase.Exist(reqCtx, username)
	if !exist {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{Message: "user not found"})
	}

	receiver, exist := userUsecase.Exist(reqCtx, req.ToUser)
	if !exist {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "receiver not found"})
	}

	if sender.Balance < uint64(req.Amount) {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "insufficient balance"})
	}

	err := transactionUsecase.Send(reqCtx, &sender, &receiver, uint64(req.Amount))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to make transaction"})
	}

	err = userUsecase.UpdateBalance(reqCtx, sender.ID, -uint64(req.Amount))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to update sender's balance"})
	}

	err = userUsecase.UpdateBalance(reqCtx, receiver.ID, uint64(req.Amount))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to update receiver's balance"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "coins sent successfully"})
}
