package handler

import (
	"log/slog"
	"store-review/internal/config"
	"store-review/internal/handler/login"
	"store-review/internal/handler/middleware"
	"store-review/internal/handler/register"
	"store-review/internal/handler/store"
	"store-review/internal/infrastructure/gen"
	"store-review/internal/infrastructure/postgres/repository"
	"store-review/internal/usecase/auth"
	suc "store-review/internal/usecase/store"

	"github.com/gin-gonic/gin"
)

func SetupRouter(q *db.Queries, cfg *config.AppConfig, logger *slog.Logger) *gin.Engine {
	userR := repository.NewUserRepository(q)
	storeR := repository.NewStoreRepository(q)

	registerUC := auth.NewRegisterUseCase(userR)
	loginUC := auth.NewLoginUseCase(userR, cfg.JWTSecret)
	sListUC := suc.NewListUseCase(storeR)

	registerH := register.NewHandler(logger, registerUC)
	loginH := login.NewHandler(logger, loginUC)
	storeH := store.NewHandler(sListUC)

	r := gin.Default()

	r.Use(middleware.LogContextMiddleware(logger), middleware.CorsMiddleware(), middleware.CSRFMiddleware(logger))
	jwt := r.Group("", middleware.JwtMiddleware(cfg.JWTSecret))

	r.POST("/register", registerH.RegisterUser)
	r.POST("/login", loginH.Login)
	jwt.GET("/stores", storeH.GetList)

	return r
}
