// Code generated by sqlc. DO NOT EDIT.
// source: author.sql

package models

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO library.authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING id, name, bio
`

type CreateAuthorParams struct {
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}

type CreateAuthorRow struct {
	ID   int64          `json:"id"`
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (CreateAuthorRow, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)
	var i CreateAuthorRow
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM library.authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio FROM library.authors
WHERE id = $1
`

type GetAuthorRow struct {
	ID   int64          `json:"id"`
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}

func (q *Queries) GetAuthor(ctx context.Context, id int64) (GetAuthorRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i GetAuthorRow
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio FROM library.authors
ORDER BY name
`

type ListAuthorsRow struct {
	ID   int64          `json:"id"`
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}

func (q *Queries) ListAuthors(ctx context.Context) ([]ListAuthorsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAuthorsRow
	for rows.Next() {
		var i ListAuthorsRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}