-- +goose Up
CREATE TABLE social_feed_followed (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    social_feed_id UUID NOT NULL REFERENCES social_feeds(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (user_id, social_feed_id)
);

-- +goose Down
DROP TABLE social_feed_followed;