package register

import (
	"log"
	"net/http"
	"store-review/internal/usecase/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	registerUseCase *auth.RegisterUseCase
}

func NewHandler(registerUseCase *auth.RegisterUseCase) *Handler {
	return &Handler{
		registerUseCase: registerUseCase,
	}
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Register request: nickname=%s, email=%s", req.Nickname, req.Email)

	input := auth.RegisterInput{
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.registerUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		log.Printf("Register usecase error: %v", err)
		// メールアドレス重複エラーの処理
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := Response{
		ID:        output.User.ID,
		Nickname:  output.User.Nickname,
		Email:     output.User.Email,
		CreatedAt: output.User.CreatedAt.String(),
		UpdatedAt: output.User.UpdatedAt.String(),
	}
	c.JSON(http.StatusCreated, res)
}
