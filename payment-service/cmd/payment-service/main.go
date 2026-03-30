package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"payment-service/internal/app"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dbURL := env("PAYMENT_DB_DSN", "postgres://postgres:postgres@localhost:5434/payment_db?sslmode=disable")
	port := env("PORT", "8082")

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	router := app.NewRouter(app.RouterDeps{DB: db})
	log.Printf("payment service started on :%s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
