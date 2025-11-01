-- name: CreateUser :one
INSERT INTO users(name, email, phone_number, password) VALUES ($1, $2, $3, $4)
RETURNING id, name, email, phone_number, created_at, updated_at;

-- name: GetUser :one
SELECT id, name, email, phone_number, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = coalesce(sqlc.narg('name'), name),
    email = coalesce(sqlc.narg('email'), email),
    phone_number = coalesce(sqlc.narg('phone_number'), phone_number),
    updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg(id)
RETURNING id, name, email, phone_number, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
