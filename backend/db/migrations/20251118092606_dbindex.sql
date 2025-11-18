-- +goose Up
-- +goose StatementBegin
-- users
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- user_sessions
CREATE INDEX IF NOT EXISTS idx_user_sessions_user_id ON user_sessions(user_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_sessions_token ON user_sessions(token);

-- agents
CREATE INDEX IF NOT EXISTS idx_agents_user_id ON agents(user_id);
CREATE INDEX IF NOT EXISTS idx_agents_status ON agents(status);
CREATE INDEX IF NOT EXISTS idx_agents_last_seen ON agents(last_seen);

CREATE INDEX IF NOT EXISTS idx_agents_user_status ON agents(user_id, status);

-- alerts
CREATE INDEX IF NOT EXISTS idx_alerts_agent_created ON alerts(agent_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_alerts_status ON alerts(status);
CREATE INDEX IF NOT EXISTS idx_alerts_severity ON alerts(severity);
CREATE INDEX IF NOT EXISTS idx_alerts_agent_status ON alerts(agent_id, status);

-- agent_queue
CREATE INDEX IF NOT EXISTS idx_agent_queue_agent_id ON agent_queue(agent_id);

-- notifications
CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications(user_id);

-- commands
CREATE INDEX IF NOT EXISTS idx_commands_agent_status ON commands(agent_id, status);
CREATE INDEX IF NOT EXISTS idx_commands_user_id ON commands(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- users
DROP INDEX IF EXISTS idx_users_email;

-- user_sessions
DROP INDEX IF EXISTS idx_user_sessions_user_id;
DROP INDEX IF EXISTS idx_user_sessions_token;

-- agents
DROP INDEX IF EXISTS idx_agents_user_id;
DROP INDEX IF EXISTS idx_agents_status;
DROP INDEX IF EXISTS idx_agents_last_seen;
DROP INDEX IF EXISTS idx_agents_user_status;

-- alerts
DROP INDEX IF EXISTS idx_alerts_agent_created;
DROP INDEX IF EXISTS idx_alerts_status;
DROP INDEX IF EXISTS idx_alerts_severity;
DROP INDEX IF EXISTS idx_alerts_agent_status;

-- agent_queue
DROP INDEX IF EXISTS idx_agent_queue_agent_id;

-- notifications
DROP INDEX IF EXISTS idx_notifications_user_id;

-- commands
DROP INDEX IF EXISTS idx_commands_agent_status;
DROP INDEX IF EXISTS idx_commands_user_id;
-- +goose StatementEnd
