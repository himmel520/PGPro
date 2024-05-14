package postgres

import (
	"context"

	"github.com/himmel520/pgPro/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository manages interactions with a PostgreSQL database.
type Repository struct {
	DB *pgxpool.Pool
}

// New creates a new instance of Repository and establishes a connection to the PostgreSQL.
func New(cfg *config.Config) (*Repository, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.GetDatabaseUrl())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Repository{DB: pool}, nil
}
