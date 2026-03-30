package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"order-service/internal/app"
)

func main() {
	dbURL := env("ORDER_DB_DSN", "postgres://postgres:postgres@localhost:5433/order_db?sslmode=disable")
	paymentURL := env("PAYMENT_SERVICE_URL", "http://localhost:8082")
	port := env("PORT", "8081")

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	router := app.NewRouter(app.RouterDeps{DB: db, Config: app.Config{PaymentBaseURL: paymentURL}})
	log.Printf("order service started on :%s", port)
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
