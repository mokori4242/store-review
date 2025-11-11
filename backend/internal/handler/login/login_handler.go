package login

import (
	"net/http"
	"store-review/internal/usecase/auth"
	"strings"

	"github.com/gin-gonic/gin"
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

	// ユースケースの入力を作成
	input := auth.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	// ユースケースを実行
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

	// JWTをHTTPOnly + Secure Cookieに設定
	c.SetCookie(
		"accessToken",
		output.AccessToken,
		24*60*60,
		"/",
		"",
		false, // 本番ではtrueに
		true,
	)

	c.JSON(http.StatusOK, gin.H{})
}
