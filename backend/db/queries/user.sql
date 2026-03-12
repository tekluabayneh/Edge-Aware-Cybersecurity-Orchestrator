-- name: GetUser :one
SELECT * from users WHERE email = $1;

-- name: GetUserByEmail :one
SELECT * from users WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (name, email, password, photo) VALUES ($1, $2, $3, $4);

-- name: UpdateUserProfile :exec
UPDATE users 
SET
    photo = COALESCE($1, photo),
    phone = COALESCE($2, phone),
    two_fa = COALESCE($3, two_fa),
    notification = COALESCE($4, notification),
    alert_notification  = COALESCE($5, alert_notification)
WHERE email = $6;


