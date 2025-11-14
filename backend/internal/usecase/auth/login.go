package auth

import (
	"context"
	"errors"
	"store-review/internal/domain/user"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	AccessToken string
	User        *user.User
}

type LoginUseCase struct {
	userRepo  user.Repository
	jwtSecret []byte
}

// NewLoginUseCase ログインユースケースを作成
func NewLoginUseCase(userRepo user.Repository, jwtSecret []byte) *LoginUseCase {
	return &LoginUseCase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// Execute ログインを実行
func (uc *LoginUseCase) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	u, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("invalid email or password")
	}

	// パスワードを検証
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := uc.generateToken(u.ID)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		AccessToken: token,
		User:        u,
	}, nil
}

// generateToken JWTトークンを生成
func (uc *LoginUseCase) generateToken(userID int64) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "store-review-backend",
			Subject:   strconv.FormatInt(userID, 10),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(uc.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
