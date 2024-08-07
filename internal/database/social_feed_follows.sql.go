// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: social_feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSocialFeedFollow = `-- name: CreateSocialFeedFollow :one
INSERT INTO social_feed_followed (id, user_id, social_feed_id, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, social_feed_id, created_at, updated_at
`

type CreateSocialFeedFollowParams struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	SocialFeedID uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (q *Queries) CreateSocialFeedFollow(ctx context.Context, arg CreateSocialFeedFollowParams) (SocialFeedFollowed, error) {
	row := q.db.QueryRowContext(ctx, createSocialFeedFollow,
		arg.ID,
		arg.UserID,
		arg.SocialFeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i SocialFeedFollowed
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SocialFeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSocialFeedFollow = `-- name: DeleteSocialFeedFollow :exec
DELETE FROM social_feed_followed WHERE id = $1 AND user_id = $2
`

type DeleteSocialFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteSocialFeedFollow(ctx context.Context, arg DeleteSocialFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteSocialFeedFollow, arg.ID, arg.UserID)
	return err
}

const getSocialFeedFollows = `-- name: GetSocialFeedFollows :many
SELECT id, user_id, social_feed_id, created_at, updated_at FROM social_feed_followed WHERE user_id = $1
`

func (q *Queries) GetSocialFeedFollows(ctx context.Context, userID uuid.UUID) ([]SocialFeedFollowed, error) {
	rows, err := q.db.QueryContext(ctx, getSocialFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SocialFeedFollowed
	for rows.Next() {
		var i SocialFeedFollowed
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SocialFeedID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
