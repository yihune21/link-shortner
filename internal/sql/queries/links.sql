-- name: CreateLink :one
INSERT INTO links (id , link , gen_key , created_at) VALUES ($1 , $2, $3 ,$4)
RETURNING *;

-- name: GetLinkById :one
SELECT * FROM links WHERE id = $1;


-- name: ListLinks :many
SELECT * FROM links;

-- name: ListLinksByLink :one
SELECT * FROM links WHERE gen_key = $1;

-- name: UpdateLinkUniqueKey :one
UPDATE links SET gen_key = $1 WHERE link = $2
RETURNING *;