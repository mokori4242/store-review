package db

import (
	"context"
	"fmt"
	"log/slog"
	"store-review/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func ConnectDB(cfgd *config.DBConfig, logger *slog.Logger) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfgd.User, cfgd.Password, cfgd.Host, cfgd.Port, cfgd.DBName, cfgd.SSLMode,
	)

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error("Failed to open DB", "error", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		logger.Error("Failed to connect DB", "error", err)
	}

	logger.Info("Connected to PostgreSQL")
	return conn
}
