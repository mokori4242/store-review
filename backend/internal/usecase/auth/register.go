package auth

import (
	"context"
	"errors"
	"store-review/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Nickname string
	Email    string
	Password string
}

type RegisterOutput struct {
	User *user.User
}

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

	newUser := user.NewUser(input.Nickname, input.Email, string(hashedPassword))

	createdUser, err := uc.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &RegisterOutput{
		User: createdUser,
	}, nil
}
