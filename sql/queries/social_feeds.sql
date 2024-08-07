-- name: CreateSocialFeed :one
INSERT INTO social_feeds (id, name, url, user_id, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllSocialFeed :many
SELECT * FROM social_feeds;