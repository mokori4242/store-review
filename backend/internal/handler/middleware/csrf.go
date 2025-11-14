package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CSRFMiddleware(logger *slog.Logger) gin.HandlerFunc {
	// CSRF保護の設定
	csrfProtection := http.NewCrossOriginProtection()

	// 信頼オリジンの設定
	err := csrfProtection.AddTrustedOrigin("http://localhost:3000")
	if err != nil {
		logger.Error("not add origin", "error", err)
	}

	// カスタムdenyハンドラの設定
	denyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Error("Origin Blocked", "origin", r.Header.Get("Origin"), "path", r.URL.Path)
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
	csrfProtection.SetDenyHandler(denyHandler)

	return func(c *gin.Context) {
		csrfProtection.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	}
}
