package main

import (
	"store-review/internal/config"
	cfgdb "store-review/internal/config/db"
	cfglog "store-review/internal/config/log"
	"store-review/internal/handler"
	"store-review/internal/infrastructure/gen"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	logger := cfglog.NewSlog()

	conn := cfgdb.ConnectDB(cfg.DB, logger)
	q := db.New(conn)

	r := handler.SetupRouter(q, cfg, logger)
	err := r.Run()
	if err != nil {
		logger.Error("Failed to run server", "error", err)
	}
}
