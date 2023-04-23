-- name: GetActivityGroup :one
SELECT * FROM activities
WHERE id = ? LIMIT 1;

-- name: ListActivityGroups :many
SELECT * FROM activities ORDER BY id DESC;

-- name: DeleteActivityGroup :exec
DELETE FROM activities
WHERE id = ?;

-- name: UpdateActivityGroup :exec
UPDATE activities
SET title = ?
WHERE id = ?;

-- name: CreateActivityGroup :execresult
INSERT INTO activities (title, email)
VALUES (?, ?);
