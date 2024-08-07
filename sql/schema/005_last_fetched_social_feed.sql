-- +goose Up
ALTER TABLE social_feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE social_feeds DROP COLUMN last_fetched_at;