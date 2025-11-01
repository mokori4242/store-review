-- name: CreateUserCar :one
INSERT INTO user_cars (user_id, registration_number, model) VALUES ($1, $2, $3)
RETURNING id, user_id, registration_number, model, created_at;

-- name: GetUserCar :one
SELECT id, user_id, registration_number, model, created_at
FROM user_cars
WHERE id = $1;

-- name: GetUserCarByRegistrationNumberAndModel :one
SELECT id, user_id, registration_number, model, created_at
FROM user_cars
WHERE registration_number = $1 AND model = $2;

-- name: GetUserCarsByUserId :many
SELECT id, user_id, registration_number, model, created_at
FROM user_cars
WHERE user_id = $1;

-- name: UpdateUserCar :one
UPDATE user_cars SET user_id = $1, registration_number = $2, model = $3
WHERE id = $4 RETURNING id, user_id, registration_number, model, created_at;

-- name: DeleteUserCar :exec
DELETE FROM user_cars WHERE id = $1;
