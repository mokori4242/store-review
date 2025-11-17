package log

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

type Handler struct {
	slog.Handler
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value("requestId").(string); ok {
		r.AddAttrs(slog.String("requestId", requestID))
	}
	if userID, ok := ctx.Value("userId").(string); ok {
		r.AddAttrs(slog.String("userId", userID))
	}
	return h.Handler.Handle(ctx, r)
}

func NewSlog() *slog.Logger {
	baseHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
	customHandler := &Handler{Handler: baseHandler}
	return slog.New(customHandler)
}
