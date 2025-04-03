-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "auth_user" (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "auth_user";
-- +goose StatementEnd
