-- +goose Up

CREATE TABLE links(
 id UUID PRIMARY KEY,
 short_link TEXT UNIQUE NOT NULL,
 original_link TEXT UNIQUE NOT NULL,
 user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
 created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE links;