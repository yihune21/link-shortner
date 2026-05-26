-- +goose Up

CREATE TABLE users(
 id UUID PRIMARY KEY,
 name TEXT UNIQUE NOT NULL,
 email TEXT UNIQUE,
 password TEXT UNIQUE,
 is_verified BOOLEAN DEFAULT FALSE,
 created_at TIMESTAMP NOT NULL , 
 updated_At TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;