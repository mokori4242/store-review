package middleware

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LogContextMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// userIdも渡せると良さそう
		requestId := uuid.NewString()
		ctx := context.WithValue(c.Request.Context(), "requestId", requestId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		logger.InfoContext(ctx, c.FullPath())
	}
}
