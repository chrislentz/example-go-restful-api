--------------------------
-- User Queries --
--------------------------

-- name: GetUserByID :one
SELECT * FROM users WHERE uuid = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at ASC;
