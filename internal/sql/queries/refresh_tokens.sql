-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (id , exp_date,token ,user_id , created_at) VALUES ($1 , $2, $3 ,$4,$5)
RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE user_id=$1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens WHERE user_id=$1;
