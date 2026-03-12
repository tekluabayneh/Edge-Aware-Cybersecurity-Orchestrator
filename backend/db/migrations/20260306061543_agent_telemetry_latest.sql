-- +goose Up
-- +goose StatementBegin
CREATE TABLE agent_telemetry_latest (
    agent_id          TEXT PRIMARY KEY,               
    machine_id        TEXT NOT NULL,                   
    agent_token       TEXT,                             
    last_updated      TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    system_data       JSONB,
    security_data     JSONB,
    processes_data    JSONB,
    integrity_data    JSONB,
    network_data      JSONB,
    expires_at        TIMESTAMPTZ                       
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop TABLE agent_telemetry_latest;
-- +goose StatementEnd
