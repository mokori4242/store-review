package auth

import (
	"context"
	"errors"
	"go-gin/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

// RegisterInput ユーザー登録の入力
type RegisterInput struct {
	Nickname string
	Email    string
	Password string
}

// RegisterOutput ユーザー登録の出力
type RegisterOutput struct {
	User *user.User
}

// RegisterUseCase ユーザー登録のユースケース
type RegisterUseCase struct {
	userRepo user.Repository
}

// NewRegisterUseCase ユーザー登録ユースケースを作成
func NewRegisterUseCase(userRepo user.Repository) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo: userRepo,
	}
}

// Execute ユーザー登録を実行
func (uc *RegisterUseCase) Execute(ctx context.Context, input RegisterInput) (*RegisterOutput, error) {
	// メールアドレスの重複チェック（ビジネスルールバリデーション）
	existingUser, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// ユーザードメインモデルを作成
	newUser := user.NewUser(input.Nickname, input.Email, string(hashedPassword))

	// ユーザーを永続化
	createdUser, err := uc.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &RegisterOutput{
		User: createdUser,
	}, nil
}
