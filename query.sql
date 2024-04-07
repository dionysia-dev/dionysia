-- name: CreateInput :one
INSERT INTO inputs (id, name, format) VALUES ($1, $2, $3) RETURNING *;

-- name: GetInput :one
SELECT * FROM inputs WHERE id = $1;

-- name: DeleteInput :exec
DELETE FROM inputs WHERE id = $1;
