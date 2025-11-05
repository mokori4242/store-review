package handler

import (
	"go-gin/internal/config"
	cfgdb "go-gin/internal/config/db"
	"go-gin/internal/handler/register"
	"go-gin/internal/infrastructure/gen"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	cfg := config.Load()

	conn := cfgdb.ConnectDB(cfg.DB)
	q := db.New(conn)
	re := register.NewHandler(q)

	r := gin.Default()

	r.POST("/register", re.RegisterUser)

	return r
}
