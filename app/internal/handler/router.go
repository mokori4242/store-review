package handler

import (
	"go-gin/internal/config"
	"go-gin/internal/handler/login"
	"go-gin/internal/handler/register"
	"go-gin/internal/infrastructure/gen"

	"github.com/gin-gonic/gin"
)

func SetupRouter(q *db.Queries, cfg *config.AppConfig) *gin.Engine {
	re := register.NewHandler(q)
	l := login.NewHandler(q, cfg.JWTSecret)

	r := gin.Default()

	r.POST("/register", re.RegisterUser)
	r.POST("/login", l.Login)

	return r
}
