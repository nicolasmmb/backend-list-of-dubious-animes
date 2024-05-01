package postgresql

import (
	"context"
	"log"
	"os"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func Connect() {
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	DB_URI := "postgres://" + DB_USER + ":" + DB_PASS + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME

	cfg, err := pgxpool.ParseConfig(DB_URI)
	if err != nil {
		log.Fatal(err)
	}
	cfg.MaxConns = 32
	cfg.MinConns = 8
	cfg.ConnConfig.Tracer = otelpgx.NewTracer()

	connection, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	dbPool = connection
}

func GetConnection() *pgxpool.Pool {
	if dbPool == nil {
		log.Fatal("Database connection is not initialized")
	}
	return dbPool
}
