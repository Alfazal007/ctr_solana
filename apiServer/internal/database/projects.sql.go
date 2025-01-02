// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: projects.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createProject = `-- name: CreateProject :one
insert into project
    (id, name, creator_id)
        values ($1, $2, $3) returning id, name, creator_id
`

type CreateProjectParams struct {
	ID        uuid.UUID
	Name      string
	CreatorID uuid.NullUUID
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, arg.ID, arg.Name, arg.CreatorID)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.CreatorID)
	return i, err
}

const getExistingProject = `-- name: GetExistingProject :one
select id, name, creator_id from project
	where name=$1
`

func (q *Queries) GetExistingProject(ctx context.Context, name string) (Project, error) {
	row := q.db.QueryRowContext(ctx, getExistingProject, name)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.CreatorID)
	return i, err
}