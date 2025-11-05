package handler

import (
	"go-gin/internal/config"
	"go-gin/internal/handler/login"
	"go-gin/internal/handler/register"
	"go-gin/internal/infrastructure/gen"
	"go-gin/internal/infrastructure/postgres/repository"
	"go-gin/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter(q *db.Queries, cfg *config.AppConfig) *gin.Engine {
	ur := repository.NewUserRepository(q)

	reuc := auth.NewRegisterUseCase(ur)
	luc := auth.NewLoginUseCase(ur, cfg.JWTSecret)

	re := register.NewHandler(reuc)
	l := login.NewHandler(luc)

	r := gin.Default()

	r.POST("/register", re.RegisterUser)
	r.POST("/login", l.Login)

	return r
}
