-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * from users WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (name, email, photo, phone, password, created_at) VALUES ($1, $2, $3, $4, $5, $6);
