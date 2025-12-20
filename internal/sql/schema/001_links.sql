-- +goose Up

CREATE TABLE links(
 id UUID PRIMARY KEY,
 link TEXT UNIQUE NOT NULL,
 gen_key TEXT UNIQUE NOT NULL,
 created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE links;