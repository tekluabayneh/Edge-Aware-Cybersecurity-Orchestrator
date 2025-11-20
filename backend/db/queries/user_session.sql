-- name: GetAllUserSession:one
SELECT * FROM user_sessions LIMIT = 50;

-- name: GetUserSessionById:one
SELECT * from user_sessions WHERE id = $1;

-- name: CreateUseSession:exec
INSERT INTO user_sessions(user_id, token, expires_at) VALUES ($1, $2, $3)

-- name: UpdateUserSessionById:exec
UPDATE user_sessions
SET
    user_id = COALESCE($1, user_id),
    token = COALESCE($2, token),
    expires_at = COALESCE($3, expires_at),
WHERE id = $4;