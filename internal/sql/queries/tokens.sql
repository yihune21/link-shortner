-- name: CreateUser :one
INSERT INTO tokens (id , exp_date ,user_id ,  created_at) VALUES ($1 , $2, $3 ,$4)
RETURNING *;

-- name: Listtokens :many
SELECT * FROM tokens;

-- name: GetToken :one
SELECT * FROM tokens WHERE email=$1;

-- name: DeleteToken :exec
DELETE FROM tokens WHERE id=$1;
