-------------------
-- Count Queries --
-------------------

-- name: CountUsers :one
SELECT count(*) FROM users;

------------------
-- Seed Queries --
------------------

-- name: SeedUser :exec
INSERT INTO users (id, uuid, name, github, twitter, mastodon) VALUES ($1, $2, $3, $4, $5, $6);
