-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS alerts (
  id BIGSERIAL PRIMARY KEY,
  agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
  agent_token TEXT NOT NULL,
  alert_type TEXT NOT NULL,
  severity TEXT NOT NULL,
  message TEXT,
  raw_payload JSONB,
  status TEXT NOT NULL DEFAULT 'new',
  risk_level TEXT,
  summary TEXT,
  performance JSONB,
  network JSONB,
  security JSONB,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE alerts;
-- +goose StatementEnd

