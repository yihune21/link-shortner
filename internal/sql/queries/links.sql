-- name: CreateLink :one
INSERT INTO links (id , link , created_at) VALUES ($1 , $2, $3)
RETURNING *;

-- name: GetLinkById :one
SELECT * FROM links WHERE id = $1;


-- name: ListLinks :many
SELECT * FROM links;

-- name: ListLinksByLink :many
SELECT * FROM links WHERE link = $1;