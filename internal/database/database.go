package database

import (
	"context"
	"fmt"
	"log/slog"
	"mastery-project/internal/config"
	"net"
	"net/url"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.Config) (*Database, error) {
	hostPort := net.JoinHostPort(cfg.Database.DBHost, strconv.Itoa(cfg.Database.DBPort))

	encodedPassword := url.QueryEscape(cfg.Database.DBPass)

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.Database.DBUser,
		encodedPassword,
		hostPort,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database connection configuration: %s", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	err2 := pool.Ping(context.Background())
	if err2 != nil {
		return nil, fmt.Errorf("failed to ping database: %s", err)
	}
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() error {
	slog.Info("database closing")
	db.Pool.Close()
	return nil
}
