-- name: Create2FA :exec
INSERT INTO user_2fa_tokens (user_id, fa_secret, isEnabled) 
VALUES ($1, $2, $3);

-- name: Get2FAByUser :one
SELECT * FROM user_2fa_tokens WHERE user_id = $1;

-- name: Is2FAEnabled :one
SELECT isEnabled, firstTime FROM user_2fa_tokens WHERE user_id = $1;

-- name: UpdateIs2FAEnabled :exec
UPDATE user_2fa_tokens SET isEnabled = $1 WHERE user_id = $2;
