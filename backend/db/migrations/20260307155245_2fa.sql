-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_2fa_tokens (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    fa_secret VARCHAR(255) NOT NULL,
    isEnabled BOOLEAN DEFAULT false,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_2fa_tokens 
-- +goose StatementEnd
