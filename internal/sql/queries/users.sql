-- name: CreateUser :one
INSERT INTO users (id , name ,email ,password ,  created_at) VALUES ($1 , $2, $3 ,$4,$5)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: Listusers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE email=$1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;

-- name: UpdateUserName :one
UPDATE users 
SET name = $1
WHERE id =$2
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users 
SET password = $1
WHERE id =$2
RETURNING *;