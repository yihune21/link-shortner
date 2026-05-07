-- +goose Up

CREATE TABLE links(
 id BIGSERIAL PRIMARY KEY,
 short_link TEXT UNIQUE NOT NULL,
 original_link TEXT UNIQUE NOT NULL,
 user_id BIGSERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
 created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE links;