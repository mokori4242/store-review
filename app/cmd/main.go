package main

import (
	"go-gin/internal/config"
	cfgdb "go-gin/internal/config/db"
	"go-gin/internal/handler"
	"go-gin/internal/infrastructure/gen"
	"log"

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
