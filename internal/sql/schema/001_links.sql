-- +goose Up

CREATE TABLE links(
 id UUID PRIMARY KEY,
 link TEXT UNIQUE NOT NULL,
 created_at TIMESTAMP
);

-- +goose Down
DROP TABLE links;