package handlers

import (
	"avito-shop/internal/auth"
	"avito-shop/internal/models"
	"avito-shop/internal/repository/user"
	"avito-shop/internal/usecase"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func AuthHandler(ctx echo.Context) error {
	db := ctx.Request().Context().Value("db").(*sqlx.DB)
	var req AuthRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request data"})
	}

	reqCtx := ctx.Request().Context()
	userRepo := user.NewRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	user_, exist := userUsecase.Exist(reqCtx, req.Username)
	if !exist {
		hash, err := auth.HashPassword(req.Password)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create user"})
		}
		user_ = models.User{
			Id:      uuid.New(),
			Name:    req.Username,
			Hash:    hash,
			Balance: 1000,
		}
		err = userUsecase.CreateUser(reqCtx, user_)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create user"})
		}
	}

	if !auth.CheckPasswordHash(req.Password, user_.Hash) {
		return ctx.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid credentials"})
	}

	token, err := auth.GenerateToken(user_.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate token"})
	}

	return ctx.JSON(http.StatusOK, AuthResponse{Token: token})
}
