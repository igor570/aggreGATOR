-- name: CreatePost :one
INSERT INTO posts (id, feed_id, title, url, description, published_at, created_at, updated_at)
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

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
INNER JOIN feeds_follow ON feeds_follow.feed_id = posts.feed_id
WHERE feeds_follow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;