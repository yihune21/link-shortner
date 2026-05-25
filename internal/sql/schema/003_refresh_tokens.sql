-- +goose Up

CREATE TABLE refresh_tokens(
 id UUID PRIMARY KEY,
 exp_date TIMESTAMP NOT NULL,
 token TEXT NOT NULL UNIQUE,
 user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
 created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE refresh_tokens;