-- name: GetAllAgentQueue :one
SELECT * FROM agent_queue LIMIT  50;

-- name: GetAgentQueueByID :one
SELECT * from agent_queue WHERE id = $1;

-- name: CreateAgentQueue :exec
INSERT INTO agent_queue(agent_id, event) VALUES ($1, $2);

-- name: UpdateAgentById :exec
UPDATE agent_queue
SET
    agent_id = COALESCE($1, agent_id),
    event = COALESCE($2, event)
WHERE id = $3;
