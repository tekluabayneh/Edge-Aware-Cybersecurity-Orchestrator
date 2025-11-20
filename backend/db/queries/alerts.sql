-- name: GetAllAgent:one
SELECT * FROM alerts LIMIT =50;

-- name: GetAgentById:one
SELECT * from alerts WHERE id = $1;

-- name: CreateAgent:exec
INSERT INTO alerts(agent_id, alert_type, severity,  status, message, raw_payload) VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateSingleAgent:exec
UPDATE alerts
SET
    agent_id = COALESCE($1, agent_id),
    alert_type = COALESCE($2, alert_type),
    severity = COALESCE($3, severity),
    message = COALESCE($4, message),
    status        = COALESCE($5, status),
    raw_payload     = COALESCE($6, raw_payload)
WHERE id = $7;

