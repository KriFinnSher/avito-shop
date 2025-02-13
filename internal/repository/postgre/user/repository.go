package user

import (
	"avito-shop/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgreRepository struct {
	db *sqlx.DB
}

func NewPostgreRepo(db *sqlx.DB) *PostgreRepository {
	return &PostgreRepository{db: db}
}

func (r *PostgreRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM Users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Name, &user.Balance, &user.Hash)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *PostgreRepository) GetUserByName(ctx context.Context, name string) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM Users WHERE username = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&user.ID, &user.Name, &user.Balance, &user.Hash)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *PostgreRepository) CreateUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO Users(id, username, balance, password_hash) VALUES($1, $2, 1000, $3)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Hash)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreRepository) UpdateUserBalance(ctx context.Context, userID uuid.UUID, amount int) error {
	query := `UPDATE Users SET balance = balance + $2 WHERE id = $1 AND balance + $2 >= 0`
	_, err := r.db.ExecContext(ctx, query, userID, amount)
	if err != nil {
		return err
	}
	return nil
}
