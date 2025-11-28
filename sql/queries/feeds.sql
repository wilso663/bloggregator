-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES(
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetAllFeeds :many
Select * FROM feeds;

-- name: GetFeedUserNameById :one
Select users.name FROM feeds
INNER JOIN users
ON feeds.user_id = users.id
Where users.id = $1 LIMIT 1;

