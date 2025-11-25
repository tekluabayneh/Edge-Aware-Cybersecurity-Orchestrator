-- name: GetAllAgent :one
SELECT * FROM agents LIMIT  50;

-- name: GetAgentById :one
SELECT * from agents WHERE id = $1;

-- name: CreateAgent :exec
INSERT INTO agents(user_id, agent_id, agent_token, machine_id, agent_version, os, status, last_seen) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetAgentByAgentId :one
SELECT * from agents WHERE agent_id = $1;

-- name: GetUserDeviceCount :one
SELECT COUNT(*) from agents WHERE user_id = $1;


-- name: UpdateSingleAgent :exec
UPDATE agents
SET
    user_id       = COALESCE($1, user_id),
    agent_id = COALESCE($2, agent_id),
    agent_token = COALESCE($3, agent_token),
    machine_id    = COALESCE($4, machine_id),
    agent_version = COALESCE($5, agent_version),
    os            = COALESCE($6, os),
    status        = COALESCE($7, status),
    last_seen     = COALESCE($8, last_seen)
WHERE id = $9;