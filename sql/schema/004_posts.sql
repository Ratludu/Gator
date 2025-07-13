-- +goose Up
CREATE TABLE posts (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	title TEXT,
	url VARCHAR(512) UNIQUE,
	description TEXT,
	published_at DATE,
	feed_id UUID NOT NULL,
	FOREIGN KEY (feed_id) REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;
