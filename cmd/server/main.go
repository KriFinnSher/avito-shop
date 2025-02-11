package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "postgres://" + getEnv("DATABASE_USER", "postgres") + ":" + getEnv("DATABASE_PASSWORD", "password") + "@" + getEnv("DATABASE_HOST", "localhost") + ":" + getEnv("DATABASE_PORT", "5432") + "/" + getEnv("DATABASE_NAME", "testdb") + "?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка пинга базы данных: ", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result string
	err := db.QueryRow("SELECT 'Hello, world!'").Scan(&result)
	if err != nil {
		http.Error(w, "Ошибка выполнения запроса к базе данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, result)
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/", handler)

	fmt.Println("Сервер запущен на порту 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
