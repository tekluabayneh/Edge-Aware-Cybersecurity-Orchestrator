-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS agent_queue (
  id BIGSERIAL PRIMARY KEY,
  agent_id BIGSERIAL NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
  event JSONB NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE agent_queue;
-- +goose StatementEnd
