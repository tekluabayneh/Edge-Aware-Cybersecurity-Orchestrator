-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_sessions (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  token TEXT NOT NULL UNIQUE,
  expires_at TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_sessions;
-- +goose StatementEnd
