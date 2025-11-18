-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS commands (
  id BIGSERIAL PRIMARY KEY,
  agent_id BIGSERIAL NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
  user_id BIGSERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  command_type TEXT NOT NULL,
  payload JSONB,
  status TEXT NOT NULL DEFAULT 'pending',
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE commands;
-- +goose StatementEnd
