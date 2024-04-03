-- name: CreateInput :one
INSERT INTO inputs (id, name, format) VALUES ($1, $2, $3) RETURNING *;
