package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"store-review/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfgd *config.DBConfig, logger *slog.Logger) *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfgd.User, cfgd.Password, cfgd.Host, cfgd.Port, cfgd.DBName, cfgd.SSLMode,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Failed to open DB", "error", err)
	}

	err = conn.Ping()
	if err != nil {
		logger.Error("Failed to connect DB", "error", err)
	}

	logger.Info("Connected to PostgreSQL")
	return conn
}
