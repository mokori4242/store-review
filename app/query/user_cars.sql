-- name: CreateUserCar :one
INSERT INTO user_cars (user_id, registration_number, model) VALUES (?, ?, ?)
RETURNING id, user_id, registration_number, model, created_at;

-- name: GetUserCar :one
SELECT id, user_id, registration_number, model, created_at
FROM user_cars
WHERE id = ?;

-- name: GetUserCarByRegistrationNumberAndModel :one
SELECT id, user_id, registration_number, model, created_at
FROM user_cars
WHERE registration_number = ? AND model = ?;

-- name: GetUserCarsByUserId :many
SELECT id, user_id, registration_number, model, created_at
FROM user_cars
WHERE user_id = ?;

-- name: UpdateUserCar :one
UPDATE user_cars SET user_id = ?, registration_number = ?, model = ?
WHERE id = ? RETURNING id, user_id, registration_number, model, created_at;

-- name: DeleteUserCar :exec
DELETE FROM user_cars WHERE id = ?;
