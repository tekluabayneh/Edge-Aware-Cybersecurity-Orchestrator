-- name: GetUserTokenById:one
SELECT * from pairing_tokens WHERE id = $1;

-- name: CreateParingToken:exec
INSERT INTO pairing_tokens(token, user_id, user_email, expires_at) VALUES ($1, $2, $3, $4);

-- name: DeleteUsedToken:exec
DELETE * from pairing_tokens WHERE user_email = $1 and user_id  = %2;