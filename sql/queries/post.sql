-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
	)
RETURNING *;

-- name: GetPostDataForUser :many
SELECT 
	posts.*,
	feed_follows.user_id
FROM 
	posts
LEFT JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE user_id = $1 
ORDER BY published_at DESC;

-- name: GetPosts :many
SELECT * 
FROM posts
LIMIT $1;
