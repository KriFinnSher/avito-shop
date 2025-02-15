package handlers

import (
	"avito-shop/internal/repository/postgre/transaction"
	"avito-shop/internal/repository/postgre/user"
	"avito-shop/internal/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InfoResponse struct {
	Coins     uint64                  `json:"coins"`
	Inventory []usecase.InventoryItem `json:"inventory"`
	History   usecase.CoinHistory     `json:"coinHistory"`
}

// InfoHandler creates response in [InfoResponse] structure for
// every registered user
func InfoHandler(ctx echo.Context) error {
	db := ctx.Request().Context().Value("db").(*sqlx.DB)
	username, ok := ctx.Get("username").(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{Message: "invalid token or missing username"})
	}

	reqCtx := ctx.Request().Context()

	userRepo := user.NewPostgreRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	transactionRepo := transaction.NewPostgreRepo(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)

	user_, exist := userUsecase.Exist(reqCtx, username)
	if !exist {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{Message: "user not found"})
	}

	coins, err := userUsecase.GetBalance(reqCtx, user_.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to get balance"})
	}

	inventory, err := transactionUsecase.GetInventory(reqCtx, user_.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to get inventory"})
	}

	history, err := transactionUsecase.GetHistory(reqCtx, user_.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "failed to get history"})
	}

	response := InfoResponse{
		Coins:     coins,
		Inventory: inventory,
		History:   history,
	}

	return ctx.JSON(http.StatusOK, response)
}
