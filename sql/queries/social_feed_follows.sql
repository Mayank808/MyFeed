-- name: CreateSocialFeedFollow :one
INSERT INTO social_feed_followed (id, user_id, social_feed_id, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSocialFeedFollows :many
SELECT * FROM social_feed_followed WHERE user_id = $1;