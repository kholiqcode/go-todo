-- name: GetActivityGroup :one
SELECT * FROM activity_groups
WHERE id = ? LIMIT 1;

-- name: ListActivityGroups :many
SELECT * FROM activity_groups ORDER BY id DESC;

-- name: DeleteActivityGroup :exec
DELETE FROM activity_groups
WHERE id = ?;

-- name: UpdateActivityGroup :exec
UPDATE activity_groups
SET title = ?
WHERE id = ?;

-- name: CreateActivityGroup :execresult
INSERT INTO activity_groups (title, email)
VALUES (?, ?);
