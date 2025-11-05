package handler

import (
	"go-gin/config"
	configdb "go-gin/config/db"
	db "go-gin/gen"
	"go-gin/handler/register"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	cfg := config.Load()

	conn := configdb.ConnectDB(cfg.DB)
	q := db.New(conn)
	re := register.NewHandler(q)

	r := gin.Default()

	r.POST("/register", re.RegisterUser)

	return r
}
