package user

import (
	"avito-shop/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// PostgreRepository is a struct type for an interaction
// with infrastructure layer (users table)
type PostgreRepository struct {
	db *sqlx.DB
}

// NewPostgreRepo returns new [PostgreRepository]
func NewPostgreRepo(db *sqlx.DB) *PostgreRepository {
	return &PostgreRepository{db: db}
}

// GetUserByID returns user with defined userID if presents, empty user otherwise
func (r *PostgreRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM Users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Name, &user.Balance, &user.Hash)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserByName returns user with defined name if presents, empty user otherwise
func (r *PostgreRepository) GetUserByName(ctx context.Context, name string) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM Users WHERE username = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&user.ID, &user.Name, &user.Balance, &user.Hash)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// CreateUser creates and add new user in users table with defined values
func (r *PostgreRepository) CreateUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO Users(id, username, balance, password_hash) VALUES($1, $2, 1000, $3)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Hash)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserBalance update user's balance with defined userID
func (r *PostgreRepository) UpdateUserBalance(ctx context.Context, userID uuid.UUID, amount int) error {
	query := `UPDATE Users SET balance = balance + $2 WHERE id = $1 AND balance + $2 >= 0`
	_, err := r.db.ExecContext(ctx, query, userID, amount)
	if err != nil {
		return err
	}
	return nil
}
