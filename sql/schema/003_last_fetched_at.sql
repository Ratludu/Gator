-- +goose Up
ALTER TABLE feeds 
ADD COLUMN last_fetched_at TIMESTAMP WITH TIME ZONE;

-- +goose Down
ALTER TABLE feeds
DROP last_fetched_at;
