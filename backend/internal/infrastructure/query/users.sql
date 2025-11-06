-- name: CreateUser :one
INSERT INTO users(nickname, email, password) VALUES ($1, $2, $3)
RETURNING id, nickname, email, created_at, updated_at;

-- name: GetUser :one
SELECT id, nickname, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET nickname = coalesce(sqlc.narg('nickname'), nickname),
    email = coalesce(sqlc.narg('email'), email),
    updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg(id)
RETURNING id, nickname, email, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, nickname, email, password, created_at, updated_at
FROM users
WHERE email = $1;
