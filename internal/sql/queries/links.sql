-- name: CreateLink :one
INSERT INTO links (id , link , created_at) VALUES ($1 , $2, $3)
RETURNING *;