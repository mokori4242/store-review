package main

import (
	"log"
	"store-review/internal/config"
	cfgdb "store-review/internal/config/db"
	"store-review/internal/handler"
	"store-review/internal/infrastructure/gen"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	conn := cfgdb.ConnectDB(cfg.DB)
	q := db.New(conn)

	r := handler.SetupRouter(q, cfg)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
