
-- name: UpsertAgentTelemetry :exec
INSERT INTO agent_telemetry_latest (
    agent_id, machine_id, agent_token,
    system_data, security_data, processes_data, integrity_data, network_data,
    last_updated 
)
VALUES (
    $1, $2, $3,
    $4::jsonb, $5::jsonb, $6::jsonb, $7::jsonb, $8::jsonb,
    NOW()
)
ON CONFLICT (agent_id) DO UPDATE SET
    agent_id       = EXCLUDED.agent_id,
    machine_id     = EXCLUDED.machine_id,
    agent_token    = EXCLUDED.agent_token,
    system_data    = EXCLUDED.system_data,
    security_data  = EXCLUDED.security_data,
    processes_data = EXCLUDED.processes_data,
    integrity_data = EXCLUDED.integrity_data,
    network_data   = EXCLUDED.network_data,
    last_updated   = NOW();

-- name: GetLatestTelemetry :one
SELECT * FROM agent_telemetry_latest WHERE agent_id = $1;

-- name: ListActiveAgents :many
SELECT agent_id, last_updated
FROM agent_telemetry_latest
WHERE last_updated > NOW() - INTERVAL '1 hour'
ORDER BY last_updated DESC;

