package login

import (
	"log/slog"
	"net/http"
	"store-review/internal/usecase/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	name   = "accessToken"
	maxAge = 24 * 60 * 60
	path   = "/"
	domain = ""
	// 本番ではtrueにする（環境変数で切り替えが良さそう）
	secure   = false
	httpOnly = true
)

type Handler struct {
	logger       *slog.Logger
	loginUseCase *auth.LoginUseCase
}

func NewHandler(logger *slog.Logger, loginUseCase *auth.LoginUseCase) *Handler {
	return &Handler{
		logger:       logger,
		loginUseCase: loginUseCase,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req Request
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.ErrorContext(ctx, "Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := auth.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.loginUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		// 認証エラーの処理
		if strings.Contains(err.Error(), "invalid email or password") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		name,
		output.AccessToken,
		maxAge,
		path,
		domain,
		secure,
		httpOnly,
	)

	c.JSON(http.StatusOK, gin.H{})
}
