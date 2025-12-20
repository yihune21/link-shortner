-- name: CreateLink :one
INSERT INTO links (id , link , short_link , created_at) VALUES ($1 , $2, $3 ,$4)
RETURNING *;

-- name: GetLinkById :one
SELECT * FROM links WHERE id = $1;


-- name: ListLinks :many
SELECT * FROM links;

-- name: ListLinksByLink :one
SELECT * FROM links WHERE short_link = $1;