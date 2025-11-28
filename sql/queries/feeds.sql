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

-- name: GetFeedByURL :one
Select * from feeds
Where feeds.url = $1 LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW()
Where id = $1;

-- name: GetNextFeedToFetch :one
Select * from feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1; 
