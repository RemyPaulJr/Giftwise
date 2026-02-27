package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	db *pgxpool.Pool
}

func startDB() (*App, error) {

	databaseUrl := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}
	// Ping db to test connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Print("Unable to ping database:", err)
	} else {
		log.Print("Ping successful.")
	}
	// Initialize App struct to pass handlers access to db connection
	app := App{db: pool}

	return &app, nil

}
