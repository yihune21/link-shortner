-- +goose Up

CREATE TABLE tokens(
 id UUID PRIMARY KEY,
 exp_date TIMESTAMP NOT NULL,
 user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
 created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE tokens;