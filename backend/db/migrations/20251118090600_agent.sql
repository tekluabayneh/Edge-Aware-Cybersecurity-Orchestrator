-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS agents (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  agent_id BIGSERIAL NOT NULL,
  agent_token TEXT NOT NULL,
  machine_id TEXT NOT NULL UNIQUE,
  agent_version TEXT,
  os TEXT,
  status TEXT,
  last_seen TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE agents;
-- +goose StatementEnd
