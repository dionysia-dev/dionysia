-- name: CreateInput :one
INSERT INTO inputs (id, name, format) VALUES ($1, $2, $3) RETURNING *;

-- name: GetInput :one
SELECT * FROM inputs WHERE id = $1;

-- name: DeleteInput :exec
DELETE FROM inputs WHERE id = $1;

-- name: CreateVideoProfile :exec
INSERT INTO video_profiles (input_id, codec, bitrate, max_key_interval, framerate, width, height)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetVideoProfiles :many
SELECT * FROM video_profiles WHERE input_id = $1;

-- name: DeleteVideoProfiles :exec
DELETE FROM video_profiles WHERE input_id = $1;

-- name: CreateAudioProfile :exec
INSERT INTO audio_profiles (input_id, rate, codec)
VALUES ($1, $2, $3);

-- name: GetAudioProfiles :many
SELECT * FROM audio_profiles WHERE input_id = $1;

-- name: DeleteAudioProfiles :exec
DELETE FROM audio_profiles WHERE input_id = $1;
