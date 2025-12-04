-- name: GetAllCommands :one
SELECT * FROM commands LIMIT 50;

-- name: GetCommandById :one
SELECT * from commands WHERE id = $1;

-- name: FetchPendingCommndByAgentId :many
SELECT * from commands WHERE agent_id = $1 AND status = 'pending';

-- name: DeleteCommandByAgentId :one
DELETE  from commands WHERE agent_id = $1 RETURNING id;


-- name: UpdateCommandStatusByAgentId :one
UPDATE commands set status = 'read' WHERE agent_id = $1 RETURNING id;


-- name: CreateCommand :exec
INSERT INTO commands(agent_id, user_id, command_type, payload, status, updated_at) VALUES ($1, $2, $3, $4, $5, $6);


-- name: UpdateCommandById :exec
UPDATE commands
SET
    agent_id = COALESCE($1, agent_id),
    user_id       = COALESCE($2, user_id),
    command_type = COALESCE($3, command_type),
    payload = COALESCE($4,  payload),
    status = COALESCE($5, status),
    updated_at = COALESCE($6, updated_at)
WHERE id = $7;


