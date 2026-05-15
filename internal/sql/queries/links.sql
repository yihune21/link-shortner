-- name: CreateLink :one
INSERT INTO links (id , short_link , original_link ,user_id ,  created_at) VALUES ($1 , $2, $3 ,$4,$5)
RETURNING *;

-- name: GetLinkById :one
SELECT * FROM links WHERE id = $1;

-- name: ListLinks :many
SELECT * FROM links;

-- name: GetLinksByUserId :many
SELECT * FROM links WHERE user_id = $1;

-- name: GetLinksByShortLink :one
SELECT * FROM links WHERE short_link = $1;

-- name: DeleteLink :exec
DELETE FROM links WHERE id = $1;