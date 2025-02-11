package db

import (
	"avito-shop/internal/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func InitDb(driverName string) *sqlx.DB {
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.Db.Host,
		config.AppConfig.Db.Port,
		config.AppConfig.Db.User,
		config.AppConfig.Db.Pass,
		config.AppConfig.Db.Name,
	)
	db, err := sqlx.Connect(driverName, dbInfo)
	if err != nil {
		return nil // TODO: add wrap on this error
	}
	return db
}

func MakeMigrations() {
	dbLine := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable",
		config.AppConfig.Db.User,
		config.AppConfig.Db.Pass,
		config.AppConfig.Db.Port,
		config.AppConfig.Db.Name,
	)
	m, err := migrate.New("file://migrations", dbLine)
	if err != nil {
		log.Fatalf("Error creating migration: %v", err)
	}

	err = m.Up() // или m.Down() для отката
	if err != nil {
		// TODO: add wrap on this error
	}
}
