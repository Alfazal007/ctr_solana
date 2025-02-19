-- name: CreateProjectImage :one
insert into project_images
    (public_id, project_id, secure_url)
        values ($1, $2, $3) returning *;

-- name: GetProjectImages :many
select * from project_images
	where project_id=$1;

-- name: GetImageByPublicId :one
select * from project_images
	where public_id=$1 and project_id=$2;

