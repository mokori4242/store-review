package user

import "time"

// User ユーザードメインモデル
type User struct {
	ID        int64
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser ユーザーを作成
func NewUser(nickname, email, password string) *User {
	now := time.Now()
	return &User{
		Nickname:  nickname,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
