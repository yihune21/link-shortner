-- +goose Up

CREATE TABLE links(
 id BIGSERIAL PRIMARY KEY,
 link TEXT UNIQUE NOT NULL,
 gen_key TEXT UNIQUE,
 created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE links;