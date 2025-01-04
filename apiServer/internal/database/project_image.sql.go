// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: project_image.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createProjectImage = `-- name: CreateProjectImage :one
insert into project_images
    (public_id, project_id, secure_url)
        values ($1, $2, $3) returning public_id, project_id, secure_url
`

type CreateProjectImageParams struct {
	PublicID  string
	ProjectID uuid.UUID
	SecureUrl string
}

func (q *Queries) CreateProjectImage(ctx context.Context, arg CreateProjectImageParams) (ProjectImage, error) {
	row := q.db.QueryRowContext(ctx, createProjectImage, arg.PublicID, arg.ProjectID, arg.SecureUrl)
	var i ProjectImage
	err := row.Scan(&i.PublicID, &i.ProjectID, &i.SecureUrl)
	return i, err
}

const getImageByPublicId = `-- name: GetImageByPublicId :one
select public_id, project_id, secure_url from project_images
	where public_id=$1 and project_id=$2
`

type GetImageByPublicIdParams struct {
	PublicID  string
	ProjectID uuid.UUID
}

func (q *Queries) GetImageByPublicId(ctx context.Context, arg GetImageByPublicIdParams) (ProjectImage, error) {
	row := q.db.QueryRowContext(ctx, getImageByPublicId, arg.PublicID, arg.ProjectID)
	var i ProjectImage
	err := row.Scan(&i.PublicID, &i.ProjectID, &i.SecureUrl)
	return i, err
}

const getProjectImages = `-- name: GetProjectImages :many
select public_id, project_id, secure_url from project_images
	where project_id=$1
`

func (q *Queries) GetProjectImages(ctx context.Context, projectID uuid.UUID) ([]ProjectImage, error) {
	rows, err := q.db.QueryContext(ctx, getProjectImages, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectImage
	for rows.Next() {
		var i ProjectImage
		if err := rows.Scan(&i.PublicID, &i.ProjectID, &i.SecureUrl); err != nil {
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
