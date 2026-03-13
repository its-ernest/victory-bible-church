package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error

	// retry loop for Docker startup
	for i := 0; i < 5; i++ {
		pool, err = pgxpool.New(ctx, connString)
		if err == nil {
			err = pool.Ping(ctx)
			if err == nil {
				log.Println("Successfully connected to Postgres!")
				return pool, nil
			}
		}

		log.Printf("Postgres not ready (attempt %d/5), retrying in 2s...", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to postgres after retries: %w", err)
}