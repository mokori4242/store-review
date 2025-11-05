package login

import (
	"database/sql"
	"errors"
	"go-gin/internal/infrastructure/gen"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	q         *db.Queries
	jwtSecret []byte
}

func NewHandler(q *db.Queries, jwtSecret []byte) *Handler {
	return &Handler{
		q:         q,
		jwtSecret: jwtSecret,
	}
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// メールアドレスでユーザーを検索
	user, err := h.q.GetUserByEmail(c.Request.Context(), req.Email)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// パスワードを検証
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// JWTトークンを生成
	token, err := GenerateToken(user.ID, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, loginResponse{AccessToken: token})
}

func GenerateToken(userID int64, jwtSecret []byte) (string, error) {
	// トークンの有効期限を設定（例：24時間）
	expirationTime := time.Now().Add(24 * time.Hour)
	// クレームを作成
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "store-review-app",
			Subject:   strconv.FormatInt(userID, 10),
		},
	}
	// トークンを生成して署名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 署名されたトークン文字列を取得
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
