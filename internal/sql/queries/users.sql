-- name: CreateUser :one
INSERT INTO users (id , name ,email ,password ,  created_at) VALUES ($1 , $2, $3 ,$4,$5)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;


-- name: Listusers :many
SELECT * FROM users;
