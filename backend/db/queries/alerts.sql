-- name: GetAllAlert :many
SELECT * FROM alerts WHERE agent_id = $1;

-- name: GetAlertById :one
SELECT * from alerts WHERE id = $1;

-- name: GetAlertByAgentId :one
SELECT * from alerts WHERE id = $1 AND agent_id = $2;

-- name: CreateAlert :exec
INSERT INTO alerts (
  agent_id,
  agent_token,
  alert_type,
  severity,
  message,
  raw_payload,
  status,
  risk_level,
  summary,
  performance,
  network,
  security,
  created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
);

-- name: UpdateSingleAlert :exec
UPDATE alerts
SET
    agent_id     = COALESCE($1, agent_id),
    agent_token  = COALESCE($2, agent_token),
    alert_type   = COALESCE($3, alert_type),
    severity     = COALESCE($4, severity),
    message      = COALESCE($5, message),
    status       = COALESCE($6, status),
    raw_payload  = COALESCE($7, raw_payload),
    created_at   = COALESCE($8, created_at),
    risk_level   = COALESCE($9, risk_level),
    summary      = COALESCE($10, summary),
    performance  = COALESCE($11, performance),
    network      = COALESCE($12, network),
    security     = COALESCE($13, security)
WHERE id = $14;