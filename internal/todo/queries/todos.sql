-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos ORDER BY id DESC;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

-- name: UpdateTodo :exec
UPDATE todos
SET title = ?
WHERE id = ?;

-- name: CreateTodo :execresult
INSERT INTO todos (title)
VALUES (?);
