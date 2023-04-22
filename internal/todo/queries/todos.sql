-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos WHERE 
CASE 
WHEN sqlc.arg(search_field) = 'activity_group_id' THEN activity_group_id 
ELSE title END =
CASE
WHEN sqlc.arg(search_field) = '' THEN ''
ELSE sqlc.arg(search_value) END
ORDER BY 
CASE WHEN sqlc.arg(sort_by) = 'title' THEN title END ASC,
CASE WHEN sqlc.arg(sort_by) = '-title' THEN title END DESC,
CASE WHEN sqlc.arg(sort_by) = 'is_active' THEN is_active END ASC,
CASE WHEN sqlc.arg(sort_by) = '-is_active' THEN is_active END DESC,
CASE WHEN sqlc.arg(sort_by) = 'createdAt' THEN created_at END ASC,
CASE WHEN sqlc.arg(sort_by) = '-createdAt' THEN created_at END DESC, 
CASE WHEN sqlc.arg(sort_by) = '' THEN activity_group_id END ASC
LIMIT ? OFFSET ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

-- name: UpdateTodo :exec
UPDATE todos
SET title = ?
WHERE id = ?;

-- name: CreateTodo :execresult
INSERT INTO todos (activity_group_id, title, is_active, priority)
VALUES (?, ?, ?, ?);
