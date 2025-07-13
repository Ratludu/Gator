
-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
	)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsFromUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id) 
    VALUES (
	    $1,
	    $2,
	    $3,
	    $4,
	    $5
		)

    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT 
	users.id as user_id,
	users.name as user_name,
	feeds.id as feed_id,
	feeds.name as feed_name
FROM
	feed_follows ff
JOIN users ON ff.user_id = users.id
JOIN feeds ON ff.feed_id = feeds.id
WHERE users.id = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;

-- name: MarkedFetched :exec
UPDATE feeds
SET
	updated_at = NOW(),
	last_fetched_at = NOW()
WHERE
	id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
