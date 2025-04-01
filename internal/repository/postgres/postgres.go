package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/indienSs/go-std/internal/config"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

func New(cfg config.PostgresConfig) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &Postgres{Db: db}, nil
}

func (p *Postgres) Close() error {
	return p.Db.Close()
}