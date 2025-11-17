package repository

import (
	"context"
	"errors"
	"store-review/internal/domain/user"
	"store-review/internal/infrastructure/gen"

	"github.com/jackc/pgx/v5"
)

// UserRepository ユーザーリポジトリの実装
type UserRepository struct {
	queries *sqlc.Queries
}

// NewUserRepository ユーザーリポジトリを作成
func NewUserRepository(queries *sqlc.Queries) user.Repository {
	return &UserRepository{
		queries: queries,
	}
}

// Create ユーザーを作成
func (r *UserRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	params := sqlc.CreateUserParams{
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: u.Password,
	}

	result, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:        result.ID,
		Nickname:  result.Nickname,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}, nil
}

// FindByEmail メールアドレスでユーザーを検索
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	result, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user.User{
		ID:        result.ID,
		Nickname:  result.Nickname,
		Email:     result.Email,
		Password:  result.Password,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}, nil
}

// FindByID IDでユーザーを検索
func (r *UserRepository) FindByID(ctx context.Context, id int64) (*user.User, error) {
	result, err := r.queries.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user.User{
		ID:        result.ID,
		Nickname:  result.Nickname,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}, nil
}
