-- name: GetAllAlert :one
SELECT * FROM alerts LIMIT 50;

-- name: GetAlertById :one
SELECT * from alerts WHERE id = $1;

-- name: CreateAlert :exec
INSERT INTO alerts (
  agent_id,
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
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
);

-- name: UpdateSingleAlert :exec
UPDATE alerts
SET
    agent_id     = COALESCE($1, agent_id),
    alert_type   = COALESCE($2, alert_type),
    severity     = COALESCE($3, severity),
    message      = COALESCE($4, message),
    status       = COALESCE($5, status),
    raw_payload  = COALESCE($6, raw_payload),
    created_at   = COALESCE($7, created_at),
    risk_level   = COALESCE($8, risk_level),
    summary      = COALESCE($9, summary),
    performance  = COALESCE($10, performance),
    network      = COALESCE($11, network),
    security     = COALESCE($12, security)
WHERE id = $13;
