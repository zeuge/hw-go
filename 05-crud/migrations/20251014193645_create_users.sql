-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('guest', 'user', 'admin');

CREATE TABLE IF NOT EXISTS users (
    id              UUID PRIMARY KEY            NOT NULL,
    name            TEXT                        NOT NULL,
    email           TEXT UNIQUE                 NOT NULL,
    role            user_role                   NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE    NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE    NOT NULL
);

CREATE INDEX IF NOT EXISTS user_name_idx ON users (name);
CREATE INDEX IF NOT EXISTS user_role_idx ON users (role);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;
-- +goose StatementEnd
