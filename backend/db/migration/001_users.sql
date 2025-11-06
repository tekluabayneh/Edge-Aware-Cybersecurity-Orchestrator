-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  photo TEXT ,
  phone TEXT,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TEXT NOT NULL,
);

-- +goose Down
DROP TABLE users;
