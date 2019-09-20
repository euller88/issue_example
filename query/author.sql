-- name: GetAuthor :one
SELECT * FROM library.authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM library.authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO library.authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM library.authors
WHERE id = $1;