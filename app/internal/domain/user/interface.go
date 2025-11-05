package user

import "context"

// Repository ユーザーリポジトリのインターフェース
type Repository interface {
	// Create ユーザーを作成
	Create(ctx context.Context, user *User) (*User, error)

	// FindByEmail メールアドレスでユーザーを検索
	FindByEmail(ctx context.Context, email string) (*User, error)

	// FindByID IDでユーザーを検索
	FindByID(ctx context.Context, id int64) (*User, error)
}
