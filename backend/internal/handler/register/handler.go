package register

import (
	"log/slog"
	"net/http"
	"store-review/internal/usecase/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger          *slog.Logger
	registerUseCase *auth.RegisterUseCase
}

func NewHandler(logger *slog.Logger, registerUseCase *auth.RegisterUseCase) *Handler {
	return &Handler{
		logger:          logger,
		registerUseCase: registerUseCase,
	}
}

func (h *Handler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.ErrorContext(ctx, "Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := auth.RegisterInput{
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.registerUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		h.logger.ErrorContext(ctx, "Register usecase error", "error", err)
		// メールアドレス重複エラーの処理
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := response{
		ID:        output.User.ID,
		Nickname:  output.User.Nickname,
		Email:     output.User.Email,
		CreatedAt: output.User.CreatedAt.String(),
		UpdatedAt: output.User.UpdatedAt.String(),
	}
	c.JSON(http.StatusCreated, res)
}
