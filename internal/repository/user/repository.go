package user

import (
	"avito-shop/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM Users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&user.Id, &user.Name, &user.Balance, &user.Hash)
	if err != nil {
		return models.User{}, err // TODO: wrap this error
	}
	return user, nil
}
func (r *Repository) GetUserByName(ctx context.Context, name string) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM Users WHERE username = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&user.Id, &user.Name, &user.Balance, &user.Hash)
	if err != nil {
		return models.User{}, err // TODO: wrap this error
	}
	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO Users(id, username, balance, password_hash) VALUES($1, $2, 1000, $3)`
	_, err := r.db.ExecContext(ctx, query, user.Id, user.Name, user.Hash)
	if err != nil {
		return err // TODO: wrap this error
	}
	return nil
}

func (r *Repository) UpdateUserBalance(ctx context.Context, userId uuid.UUID, amount int) error {
	query := `UPDATE Users SET balance = balance + $2 WHERE id = $1 AND balance + $2 >= 0`
	_, err := r.db.ExecContext(ctx, query, userId, amount)
	if err != nil {
		return err // TODO: wrap this error
	}
	return nil
}
