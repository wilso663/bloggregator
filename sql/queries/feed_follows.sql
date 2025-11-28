-- name: CreateFeedFollow :one
With inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (
    $1, $2, $3, $4, $5
  )
  RETURNING *
) SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
  FROM inserted_feed_follow
  INNER JOIN users ON users.id = user_id
  INNER JOIN feeds ON feeds.id = feed_id;

-- name: GetFeedFollowsForUser :many
Select feed_follows.*, users.name AS user_name, feeds.name AS feed_name
FROM feed_follows
INNER JOIN users on users.id = user_id
INNER JOIN feeds on feeds.id = feed_id
WHERE users.name = $1;

-- name: DeleteFeedFollowByUserAndFeed :exec
DELETE FROM feed_follows
WHERE feed_follows.feed_id = $1 AND feed_follows.user_id = $2;