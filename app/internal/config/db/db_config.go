package db

import (
	"database/sql"
	"fmt"
	"go-gin/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(cfgd *config.DBConfig) *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfgd.User, cfgd.Password, cfgd.Host, cfgd.Port, cfgd.DBName, cfgd.SSLMode,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return conn
}
