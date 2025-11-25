-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  photo TEXT ,
  phone TEXT,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE users;

-- +goose StatementEnd
