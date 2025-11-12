package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CSRFMiddleware() gin.HandlerFunc {
	// CSRF保護の設定
	csrfProtection := http.NewCrossOriginProtection()

	// 信頼オリジンの設定
	err := csrfProtection.AddTrustedOrigin("http://localhost:3000")
	if err != nil {
		log.Fatalf("fatal not origin: %v", err)
	}

	// カスタムdenyハンドラの設定
	denyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Origin Blocked: Origin=%s, Path=%s", r.Header.Get("Origin"), r.URL.Path)
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
	csrfProtection.SetDenyHandler(denyHandler)

	return func(c *gin.Context) {
		csrfProtection.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	}
}
