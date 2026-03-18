package status

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Name() string {
	return "postgres"
}

func (r *PostgresRepository) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}
