// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: projects.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const countRunningProjects = `-- name: CountRunningProjects :one
select count(*) from project 
	where started=true 
		AND completed=false
			AND creator_id=$1 limit 1
`

func (q *Queries) CountRunningProjects(ctx context.Context, creatorID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRunningProjects, creatorID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createProject = `-- name: CreateProject :one
insert into project
    (id, name, creator_id)
        values ($1, $2, $3) returning id, name, started, completed, creator_id, votes
`

type CreateProjectParams struct {
	ID        uuid.UUID
	Name      string
	CreatorID uuid.UUID
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, arg.ID, arg.Name, arg.CreatorID)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Started,
		&i.Completed,
		&i.CreatorID,
		&i.Votes,
	)
	return i, err
}

const endProject = `-- name: EndProject :one
update project set completed=true
	where id=$1 returning id, name, started, completed, creator_id, votes
`

func (q *Queries) EndProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRowContext(ctx, endProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Started,
		&i.Completed,
		&i.CreatorID,
		&i.Votes,
	)
	return i, err
}

const getCreatorProjects = `-- name: GetCreatorProjects :many
select id, name, started, completed, creator_id, votes from project
	where creator_id=$1
`

func (q *Queries) GetCreatorProjects(ctx context.Context, creatorID uuid.UUID) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getCreatorProjects, creatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Started,
			&i.Completed,
			&i.CreatorID,
			&i.Votes,
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

const getExistingProject = `-- name: GetExistingProject :one
select id, name, started, completed, creator_id, votes from project
	where name=$1
`

func (q *Queries) GetExistingProject(ctx context.Context, name string) (Project, error) {
	row := q.db.QueryRowContext(ctx, getExistingProject, name)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Started,
		&i.Completed,
		&i.CreatorID,
		&i.Votes,
	)
	return i, err
}

const getExistingProjectById = `-- name: GetExistingProjectById :one
select id, name, started, completed, creator_id, votes from project
	where id=$1
`

func (q *Queries) GetExistingProjectById(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRowContext(ctx, getExistingProjectById, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Started,
		&i.Completed,
		&i.CreatorID,
		&i.Votes,
	)
	return i, err
}

const increaseVoteCount = `-- name: IncreaseVoteCount :exec
update project set votes=$1
	where id=$2
`

type IncreaseVoteCountParams struct {
	Votes int32
	ID    uuid.UUID
}

func (q *Queries) IncreaseVoteCount(ctx context.Context, arg IncreaseVoteCountParams) error {
	_, err := q.db.ExecContext(ctx, increaseVoteCount, arg.Votes, arg.ID)
	return err
}

const startProject = `-- name: StartProject :one
update project set started=true
	where id=$1 returning id, name, started, completed, creator_id, votes
`

func (q *Queries) StartProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRowContext(ctx, startProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Started,
		&i.Completed,
		&i.CreatorID,
		&i.Votes,
	)
	return i, err
}
