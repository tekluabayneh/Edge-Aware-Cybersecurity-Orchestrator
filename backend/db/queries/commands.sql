
-- name: GetAllCommands:one
SELECT * FROM commands LIMIT =50;

-- name: GetCommandById:one
SELECT * from commands WHERE id = $1;

-- name: CreateCommand:exec
INSERT INTO commands(agent_id, user_id, command_type, payload, status, is_read, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)

-- name: UpdateCommandById:exec
UPDATE commands
SET
    agent_id = COALESCE($1, agent_id),
    user_id       = COALESCE($2, user_id),
    command_type = COALESCE($3, command_type),
    payload = COALESCE($4,  payload)
    status = COALESCE($5, status)
    is_read = COALESCE($6, is_read)
    updated_at = COALESCE($7j, updated_at)
WHERE id = $8;
