// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user_queries.sql

package sqlc

import (
	"context"
)

const getUserByID = `-- name: GetUserByID :one

SELECT id, uuid, name, github, twitter, mastodon, created_at, updated_at, deleted_at FROM users WHERE uuid = $1 AND deleted_at IS NULL LIMIT 1
`

// ------------------------
// User Queries --
// ------------------------
func (q *Queries) GetUserByID(ctx context.Context, uuid string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, uuid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Name,
		&i.Github,
		&i.Twitter,
		&i.Mastodon,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, uuid, name, github, twitter, mastodon, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL ORDER BY created_at ASC
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Name,
			&i.Github,
			&i.Twitter,
			&i.Mastodon,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
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