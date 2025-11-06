package handler

import (
	"store-review/internal/config"
	"store-review/internal/handler/login"
	"store-review/internal/handler/middleware"
	"store-review/internal/handler/register"
	"store-review/internal/infrastructure/gen"
	"store-review/internal/infrastructure/postgres/repository"
	"store-review/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter(q *db.Queries, cfg *config.AppConfig) *gin.Engine {
	ur := repository.NewUserRepository(q)

	reuc := auth.NewRegisterUseCase(ur)
	luc := auth.NewLoginUseCase(ur, cfg.JWTSecret)

	re := register.NewHandler(reuc)
	l := login.NewHandler(luc)

	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	r.POST("/register", re.RegisterUser)
	r.POST("/login", l.Login)

	return r
}
