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

// InitDB initialize database with AppConfig's parameters for a defined driver
func InitDB(driverName string) *sqlx.DB {
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DB.Host,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.User,
		config.AppConfig.DB.Pass,
		config.AppConfig.DB.Name,
	)
	db, err := sqlx.Connect(driverName, dbInfo)
	if err != nil {
		log.Fatalf("failed to connect to database")
		return nil
	}
	return db
}

// MakeMigrations use all *.up.sql files if up is true, and *.down.sql otherwise
func MakeMigrations(up bool) {
	dbLine := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable",
		config.AppConfig.DB.User,
		config.AppConfig.DB.Pass,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.Name,
	)
	m, err := migrate.New("file://migrations", dbLine)
	if err != nil {
		log.Fatalf("error creating migration")
	}

	if up {
		err = m.Up()
		if err != nil {
			fmt.Println(err, "(but it's fine)")
		}
	} else {
		err = m.Down()
		if err != nil {
			log.Fatalf("error rolling back migration")
		}
	}
}
