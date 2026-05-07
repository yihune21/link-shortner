-- name: CreateUser :one
INSERT INTO users (id , name ,email ,  created_at) VALUES ($1 , $2, $3 ,$4)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;


-- name: Listusers :many
SELECT * FROM users;
