package login

import (
	"net/http"
	"store-review/internal/usecase/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	name     = "accessToken"
	maxAge   = 24 * 60 * 60
	path     = "/"
	domain   = ""
	secure   = false
	httpOnly = true
)

type Handler struct {
	loginUseCase *auth.LoginUseCase
}

func NewHandler(loginUseCase *auth.LoginUseCase) *Handler {
	return &Handler{
		loginUseCase: loginUseCase,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
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
		secure, // 本番ではtrueに
		httpOnly,
	)

	c.JSON(http.StatusOK, gin.H{})
}
